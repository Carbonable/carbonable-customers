FROM golang:1.21.6-alpine3.19 as builder

WORKDIR /srv/app

RUN apk update && apk upgrade && apk add --no-cache build-base ca-certificates

# Better caching
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download

COPY . .
RUN set -eux; \
  go build -ldflags="-linkmode external -extldflags -static" -o api cmd/api/main.go; \
  go build -ldflags="-linkmode external -extldflags -static" -o migrate cmd/migrate/main.go

# Add non-root user
RUN set -eux; \
  addgroup --system carbonable; \
  adduser --system carbonable --ingroup carbonable
USER carbonable:carbonable


FROM alpine:3.19 as production

WORKDIR /srv/app

# Copy user
COPY --from=builder /etc/passwd /etc/passwd
# Copy ssl certs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /srv/app/api ./api
COPY --from=builder /srv/app/migrate ./migrate

EXPOSE 8080

USER carbonable
