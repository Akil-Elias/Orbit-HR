package main

import (
	"Orbit_HR/internal/database"
	"fmt"
	"log"
	"net/http"
	"os"

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
	database.Conn(dbKey)

	fmt.Println("running on" + port)
	log.Fatal(http.ListenAndServe(port, nil))

}
