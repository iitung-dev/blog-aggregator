package main

import (
	"context"
	"fmt"
	"strconv"
)


func handlerBrowse(s *state, cmd command) error {

	limit := 2

	if len(cmd.Args) > 0 {
		parsedLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("failed to convert the limit argument to string: %w", err)
		}
		limit = parsedLimit
	}

	getPosts, err := s.db.GetPosts(context.Background(), int32(limit))
	if err != nil {
		return fmt.Errorf("failed to get posts: %w", err)
	}

	for _, post := range getPosts {
		fmt.Printf("Title: %v\n", post.Title)
		fmt.Printf("Url: %v\n", post.Url)
		fmt.Printf("Description: %v\n", post.Description)
		fmt.Printf("Publish date: %v\n", post.PublishedAt)
	}

	return nil
}