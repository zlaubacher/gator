package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login command requires username")
	}
	username := cmd.Args[0]
	
	if err := s.config.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("username has been set to %s\n", username)

	return nil
}