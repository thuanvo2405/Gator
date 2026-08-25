package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/thuanvo2405/Gator/internal/config"
	"github.com/thuanvo2405/Gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	dbURL := cfg.DB_URL

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	dbQueries := database.New(db)

	dataState := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	commands := commands{
		cmd: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerGetUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", handlerAddFeed)
	commands.register("feeds", handlerGetFeeds)
	commands.register("follow", handlerFollow)
	commands.register("following", handlerFollowing)

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
