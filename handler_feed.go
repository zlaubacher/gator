package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("addFeed requires two arguments: \"rss name\" and \"rss url\"")
	}

	rssName := cmd.Args[0]
	rssURL := cmd.Args[1]

	currentUserName := s.config.CurrentUserName

	user, err := s.database.GetUser(context.Background(), currentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      rssName,
		Url:       rssURL,
		UserID:    user.ID,
	}

	feed, err := s.database.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	fmt.Println("Feed created successfully:")
	fmt.Printf(" * ID:       %v\n", feed.ID)
	fmt.Printf(" * Name:     %v\n", feed.Name)
	fmt.Printf(" * URL:      %v\n", feed.Url)
	fmt.Printf(" * User ID:  %v\n", feed.UserID)

	return nil
}
