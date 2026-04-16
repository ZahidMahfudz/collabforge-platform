package config

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB() *pgxpool.Pool {
	//load env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	//ambil konfigurasi dari env
	host := GetEnv("DB_HOST")
	port := GetEnv("DB_PORT")
	user := GetEnv("DB_USER")
	password := GetEnv("DB_PASSWORD")
	dbname := GetEnv("DB_NAME")
	sslmode := GetEnv("DB_SSLMODE")

	//buat connection string
	dsn := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=" + sslmode

	//buat context dengan timeout untuk koneksi database
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//koneksikan ke database
	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	//cek koneksi dengan ping
	if err := dbpool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	log.Printf("Successfully connected to database %s at %s:%s", dbname, host, port)

	return dbpool
}