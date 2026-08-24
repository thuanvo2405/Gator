package main

import (
	"context"
	"fmt"
)

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("No users found")
	}

	for _, feed := range feeds {
		fmt.Printf("* Name: %s\n", feed.FeedName)
		fmt.Printf("  Url: %s\n", feed.FeedUrl)
		fmt.Printf("  User: %s\n", feed.UserName)
		fmt.Println("=====================================")
	}
	return nil
}
