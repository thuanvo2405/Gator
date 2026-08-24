package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	baseURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), baseURL)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	fmt.Printf("%+v\n", feed)
	return nil
}
