package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	//Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	dbKey := os.Getenv("DB_key")

	//DB connection
	dbpool, err := pgxpool.New(context.Background(), dbKey)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer dbpool.Close()

	if err := dbpool.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	fmt.Println("Database connected succesfully")

	fmt.Println("running on" + port)
	log.Fatal(http.ListenAndServe(port, nil))

}
