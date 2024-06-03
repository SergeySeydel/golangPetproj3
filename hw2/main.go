package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"myproject/db"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=postgres dbname=bankstoredb sslmode=disable host=localhost port=5435"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(dbConn)
	ctx := context.Background()

	// Create Entry
	newEntry, err := queries.CreateEntry(ctx, "My first entry")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created Entry:", newEntry)

	// Get Entry
	entry, err := queries.GetEntry(ctx, newEntry.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Got Entry:", entry)

	// List Entries
	entries, err := queries.ListEntries(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("List of Entries:", entries)

	// Update Entry
	updatedEntry, err := queries.UpdateEntry(ctx, "Updated content", newEntry.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated Entry:", updatedEntry)

	// Delete Entry
	err = queries.DeleteEntry(ctx, newEntry.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Entry:", newEntry.ID)

	// Create Transfer
	newTransfer, err := queries.CreateTransfer(ctx, db.CreateTransferParams{
		FromEntryID: newEntry.ID,
		ToEntryID:   newEntry.ID,
		Amount:      100.00,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created Transfer:", newTransfer)

	// Get Transfer
	transfer, err := queries.GetTransfer(ctx, newTransfer.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Got Transfer:", transfer)

	// List Transfers
	transfers, err := queries.ListTransfers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("List of Transfers:", transfers)

	// Update Transfer
	updatedTransfer, err := queries.UpdateTransfer(ctx, db.UpdateTransferParams{
		FromEntryID: newEntry.ID,
		ToEntryID:   newEntry.ID,
		Amount:      200.00,
		ID:          newTransfer.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated Transfer:", updatedTransfer)

	// Delete Transfer
	err = queries.DeleteTransfer(ctx, newTransfer.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Transfer:", newTransfer.ID)
}
