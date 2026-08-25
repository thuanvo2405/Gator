package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/thuanvo2405/Gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	url := cmd.args[0]

	current_user, err := s.db.GetUser(context.Background(), s.cfg.Username)
	if err != nil {
		return fmt.Errorf("Failed to get user: %w", err)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %v", err)
	}

	ffRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    current_user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}

	fmt.Printf("FeedFollow created:\n")
	fmt.Printf("* User: %s\n", ffRow.UserName)
	fmt.Printf("* Feed: %s\n", ffRow.FeedName)

	return nil
}
