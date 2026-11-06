package main

import (
	"context"
	"fmt"
	"log"

	"github.com/automdev/autom-go"
)

func main() {
	client, err := autom.NewClient("", nil) // reads AUTOM_API_KEY
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	response, err := client.GoogleSearch(ctx, &autom.GoogleSearchParams{Query: "web scraping api"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(response.Body))
}
