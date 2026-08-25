package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	current_user, err := s.db.GetUser(context.Background(), s.cfg.Username)
	if err != nil {
		return fmt.Errorf("Fail fetch current user %v", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), current_user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", current_user.Name)
	for _, ff := range feedFollows {
		fmt.Printf("- '%s'\n", ff.FeedName)
	}

	return nil
}
