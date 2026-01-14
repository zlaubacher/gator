package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("follow command requires one argument: url of feed to be followed")
	}

	rssURL := cmd.Args[0]
	currentUserName := s.config.CurrentUserName

	feed, err := s.database.GetFeedWithUrl(context.Background(), rssURL)
	if err != nil {
		return fmt.Errorf("error retrieving feed using url given: %w", err)
	}

	user, err := s.database.GetUser(context.Background(), currentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.database.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Printf("Current User %s has followed the feed: %s", feedFollow.UserName, feedFollow.FeedName)

	return nil
}
