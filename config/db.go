package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func SetupDataBase() sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("erro no env a aplicação vai explodir")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectionStr := fmt.Sprintf("host=%s port%s user=%s password=%s dbname%s sslmode=disabe", dbHost, dbPort, dbUser, dbPassword, dbName)

	dbConnection, err := sql.Open("postgres", connectionStr)

	if err != nil {
		log.Fatal(err)
	}

	err = dbConnection.Ping()

	fmt.Println("essa porra conectou ")

	return *dbConnection
}
