package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/thuanvo2405/Gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get next feed: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark feed fetched: %w", err)
	}

	feedData, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	for _, item := range feedData.Channel.Item {
		pubDate := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			pubDate = sql.NullTime{Time: t, Valid: true}
		} else if t, err := time.Parse(time.RFC1123, item.PubDate); err == nil {
			pubDate = sql.NullTime{Time: t, Valid: true}
		}

		desc := sql.NullString{}
		if item.Description != "" {
			desc = sql.NullString{String: item.Description, Valid: true}
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: desc,
			PublishedAt: pubDate,
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "unique constraint") {
				continue
			}
			log.Printf("couldn't create post: %v", err)
		}
	}

	log.Printf("Feed %s fetched, %d posts processed", nextFeed.Name, len(feedData.Channel.Item))
	return nil
}
