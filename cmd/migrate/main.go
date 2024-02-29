package main

import (
	"flag"

	"github.com/carbonable/carbonable-customers/internal/customer"
	appdb "github.com/carbonable/carbonable-customers/internal/db"
	"github.com/charmbracelet/log"
)

func main() {
	fresh := flag.Bool("fresh", false, "drop all tables before migration")
	flag.Parse()
	log.Info("Starting application migration")

	db, err := appdb.GetDbConnection()
	if err != nil {
		log.Fatalf("failed to get db connection: %v", err)
		return
	}

	if *fresh {
		log.Info("Dropping all tables")
		_ = db.Migrator().DropTable(&customer.Customer{})
	}

	_ = db.AutoMigrate(&customer.Customer{})

	log.Info("Migration done !")
}
