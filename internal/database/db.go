package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Conn(key string) {
	pool, err := pgxpool.New(context.Background(), key)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	fmt.Println("Database connected succesfully")
}
