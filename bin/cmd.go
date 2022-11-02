package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mrityunjaygr8/go-reddit/ent"
)

func main() {
	fmt.Println("yo from cmd")

	client, err := ent.Open("postgres", "postgres://postgres:example@localhost:5432/go-reddit?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}
