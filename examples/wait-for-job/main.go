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
	submitted, err := client.GoogleSearch(ctx, &autom.GoogleSearchParams{Query: "web scraping api", IsAsync: autom.Bool(true)})
	if err != nil {
		log.Fatal(err)
	}
	var accepted struct {
		JobID string `json:"job_id"`
	}
	if err := submitted.Decode(&accepted); err != nil {
		log.Fatal(err)
	}
	job, err := client.WaitForJob(ctx, accepted.JobID, &autom.WaitForJobOptions{
		OnProgress: func(job map[string]any) { fmt.Println(job["status"], job["progress"]) },
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(job.Body))
}
