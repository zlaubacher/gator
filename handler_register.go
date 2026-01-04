package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("must provide name to register")
	}

	name := cmd.Args[0]
	id := uuid.New()
	now := time.Now()

	params := database.CreateUserParams{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
	}

	user, err := s.database.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	if err := s.config.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("user %s created\n", user.Name)
	fmt.Printf("user data: %+v\n", user)

	return nil
}
