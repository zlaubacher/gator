package main

import (
	"errors"
	"fmt"
	"gator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	if _, exists := c.registeredCommands[name]; exists {
		fmt.Printf("command %q already registered \n", name)
		return
	}

	if len(name) == 0 {
		fmt.Println("command name must be given")
		return
	}

	c.registeredCommands[name] = f
}