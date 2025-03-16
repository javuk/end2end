package pkg

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func InitDB() {
	// Get DATABASE_URL from Heroku environment variables
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	dsn += "?sslmode=require"

	var err error
	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	fmt.Println("Connected to the database")
}
