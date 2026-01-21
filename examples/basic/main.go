package main

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	stoyuristov "github.com/Sunsetboy/sto-yuristov-go-sdk"
)

func main() {
	baseURL := os.Getenv("API_URL")
	authToken := os.Getenv("AUTH_TOKEN")

	if baseURL == "" || authToken == "" {
		log.Fatal("API_URL and AUTH_TOKEN environment variables are required")
	}

	client := stoyuristov.NewClient(
		baseURL,
		authToken,
		stoyuristov.WithTimeout(30*time.Second),
	)

	ctx := context.Background()

	// Example: Update question categories
	err := client.UpdateQuestion(ctx, 123, &stoyuristov.QuestionUpdateRequest{
		Categories: []int{1, 2, 3},
	})
	if err != nil {
		if errors.Is(err, stoyuristov.ErrNotFound) {
			log.Println("Question not found")
		} else {
			log.Fatalf("Failed to update question: %v", err)
		}
	}

	// Example: Update question title
	err = client.UpdateQuestion(ctx, 123, &stoyuristov.QuestionUpdateRequest{
		Title: stoyuristov.String("New Title"),
		SeoH1: stoyuristov.String("SEO H1 Heading"),
	})
	if err != nil {
		log.Fatalf("Failed to update question title: %v", err)
	}

	// Example: Update lead
	err = client.UpdateLead(ctx, 456, &stoyuristov.LeadUpdateRequest{
		Categories: []int{1, 2},
		Directions: []int{10, 20},
	})
	if err != nil {
		log.Fatalf("Failed to update lead: %v", err)
	}

	// Example: Update sentiment
	err = client.UpdateSentiment(ctx, 123, "question", 0.85)
	if err != nil {
		log.Fatalf("Failed to update sentiment: %v", err)
	}

	log.Println("All updates completed successfully")
}
