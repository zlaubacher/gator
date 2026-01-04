package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login command requires username")
	}
	username := cmd.Args[0]

	user, err := s.database.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	if err := s.config.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("logged in as %s\n", user.Name)

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.database.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Println("error resetting database:", err)
		return err
	}

	fmt.Println("database successfully deleted")
	return nil
}
