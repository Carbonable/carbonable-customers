package api

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type EchoHandlerFunc func(c echo.Context) error

type AppData struct {
	Db  *gorm.DB
	Ctx echo.Context
}

func WithDb(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}
