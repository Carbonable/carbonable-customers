package db

import (
	"database/sql"
	"os"

	"github.com/charmbracelet/log"
	_ "github.com/jackc/pgx/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Get pg database connection
func GetDbConnection() (*gorm.DB, error) {
	pgsql, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Error("failed to connect to postgres")
		return nil, err
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: pgsql,
	}), &gorm.Config{})
	if err != nil {
		log.Error("failed to get gorm connection", "error", err)
		return nil, err
	}
	return db, nil
}
