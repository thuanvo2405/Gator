package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Expected the username")
	}

	username := cmd.args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("Username has been updated: %v\n", username)

	return nil
}
