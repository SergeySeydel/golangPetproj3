package main

import (
	"context"
	"fmt"
	"log"
	"m/db"

	"github.com/jackc/pgx/v5"
)

var queries *db.Queries
var ctx = context.Background()

func main() {
	fmt.Println("Starting main")
	fmt.Println("Starting db connection...")
	conn, err := pgx.Connect(ctx, "user=postgres password=postgres dbname=bankstoredb sslmode=disable host=localhost port=5435")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	queries = db.New(conn)

	queries.CreateCard(ctx, db.CreateCardParams{Time: "Hello", Title: "Ex", Client: "Sergey", User: ""})
	cards, err := queries.ListCards(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cards)
}
