package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/iitung-dev/re-blog-aggregator/internal/database"
	"github.com/lib/pq"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Could not parse given time duration for aggregation: %v\n", err)
	}

	fmt.Printf("Collecting feeds every %v\n", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	defer ticker.Stop()

	for ;; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("Failed to scrape feed: %v\n", err)
		}
	}
}

func scrapeFeeds(s *state) error {
	next_feed, err :=  s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Failed to get next feed to fetch: %v\n", err)
	}

	mark_feed, err := s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		LastFetchAt: sql.NullTime{
			Time: time.Now(),
			Valid: true,
		},
		ID: next_feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Failed to update/mark fetched feed: %v\n", err)
	}

	rssFeed, err := fetchFeed(context.Background(), mark_feed.Url)
	if err != nil {
		return fmt.Errorf("Couldn't fetch feed: %w", err)
	}

	for _, post := range rssFeed.Channel.Item {

		formatted_pub_date, parsedError := time.Parse(time.RFC1123, post.PubDate)
		if parsedError != nil {
			log.Printf("failed to parse publication date: %v", parsedError)
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),          
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(), 
			Title: post.Title,     
			Url: post.Link,         
			Description: sql.NullString{
				String: post.Description,
				Valid: post.Description != "",
			}, 
			PublishedAt: sql.NullTime{
				Time: formatted_pub_date,
				Valid: parsedError == nil,
			}, 
			FeedID: next_feed.ID,     
		})
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
				// It's a unique constraint violation—just skip it
				continue
			}
			log.Printf("Failed to save post (%v): %v\n", post.Title, err)
			continue
		}
		log.Printf("Saved new post: %v", post.Title)
	}

	return nil
}