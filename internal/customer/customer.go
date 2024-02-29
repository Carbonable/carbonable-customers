package customer

import "github.com/oklog/ulid/v2"

type Customer struct {
	WalletAddress string    `json:"wallet_address"`
	Country       string    `json:"country"`
	Network       string    `json:"network"`
	ID            ulid.ULID `json:"id"`
}

func FromHttpRequest(wallet string, country string, network string) *Customer {
	return &Customer{
		WalletAddress: wallet,
		Country:       country,
		Network:       network,
		ID:            ulid.Make(),
	}
}
