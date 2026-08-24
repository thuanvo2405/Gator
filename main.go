package main

import (
	"fmt"
	"log"

	"github.com/thuanvo2405/Gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	err = cfg.SetUser("Thuan")
	if err != nil {
		log.Fatalf("Error setting user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading updated config: %v", err)
	}

	fmt.Printf("Config: %+v\n", cfg)
}
