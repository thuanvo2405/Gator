package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/thuanvo2405/Gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	name := cmd.args[0]
	url := cmd.args[1]

	current_user, err := s.db.GetUser(context.Background(), s.cfg.Username)
	if err != nil {
		return fmt.Errorf("Failed to get user: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    current_user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    current_user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}

	fmt.Println("Feed created successfully:")
	fmt.Printf("%+v\n", feed)
	fmt.Println()
	fmt.Printf("FeedFollow created:\n* User: %s\n* Feed: %s\n", feedFollow.UserName, feedFollow.FeedName)
	return nil
}
