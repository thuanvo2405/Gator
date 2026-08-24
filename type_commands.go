package main

import (
	"errors"

	"github.com/thuanvo2405/Gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmd map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, exists := c.cmd[cmd.name]
	if !exists {
		return errors.New("Handler not available")
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmd[name] = f
}
