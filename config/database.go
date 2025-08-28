package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sqlx.DB

func ConnectDB() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("ENV file is not found")
	}

	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	dbhost := os.Getenv("DB_HOST")

	if dbuser == "" || dbpass == "" || dbhost == "" || dbport == "" || dbname == "" {
		log.Fatal("invalid value")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpass, dbhost, dbport, dbname)

	db, err := sqlx.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Database connection failed")
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping  database : %v", err)
	}

	log.Println("Database connection succesful")

	DB = db

	return db
}
