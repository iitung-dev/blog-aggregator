package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/iitung-dev/re-blog-aggregator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Failed to get feed: %v", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Failed to create feed follow: %v", err)
	}

	fmt.Printf("You are now following %v\n", feed.Url)

	return nil
}

func handlerFollowing(s *state, _ command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Failed to get feed follows for user: %v", err)
	}

	printFeedFollow(feeds)

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.Args) < 1 {
		return fmt.Errorf("<cli> unfollow <url>\n")
	}

	feedUrl := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("Unfollow failed, cannot retrieve feed id: %v\n", err)
	}

	_, err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Failed to unfollow feed: %v\n", err)
	}

	return nil
}

func printFeedFollow(feeds []database.GetFeedFollowsForUserRow) {
	for _, feed := range feeds {
		fmt.Printf("Feed Name: %v\n", feed.FeedName)
		fmt.Printf("User Name: %v\n", feed.UserName)
	}
}