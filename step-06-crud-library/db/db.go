package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func ConnectDB() *sqlx.DB {
	// Load .env file (if exists)
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Could not load .env file, using system environment variables")
	}

	dbUser := getEnv("DB_USER", "dev")
	dbPass := getEnv("DB_PASSWORD", "dev")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "mydb")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Fatal("❌ Failed to open database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("❌ Failed to ping database:", err)
	}

	log.Println("🚀 Connected to PostgreSQL!")
	return db
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
