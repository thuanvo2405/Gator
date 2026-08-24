package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thuanvo2405/Gator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	dataState := state{
		cfg: &cfg,
	}

	commands := commands{
		cmd: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Error: Not enough arguments provided")
		os.Exit(1)
	}

	cmdName := args[1]
	cmdArgs := args[2:]

	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	err = commands.run(&dataState, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
