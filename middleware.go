package main

import (
	"context"
	"fmt"

	"github.com/thuanvo2405/Gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, c command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.Username)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}
		return handler(s, c, user)
	}
}
