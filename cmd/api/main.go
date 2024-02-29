package main

import (
	"github.com/carbonable/carbonable-customers/internal/api"
	"github.com/carbonable/carbonable-customers/internal/customer"
	appdb "github.com/carbonable/carbonable-customers/internal/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	db, err := appdb.GetDbConnection()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(api.WithDb(db))

	e.POST("/customer", customer.CreateCustomer(db))
	e.Logger.Fatal(e.Start(":8080"))
}
