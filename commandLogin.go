package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	name := cmd.args[0]

	user, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("user %s does not exist", name)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User has been set!")
	return nil
}
