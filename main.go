package main

import (
	"fmt"
	"gator/internal/config"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	s := &state{config: &cfg}

	cmds := commands{registeredCommands: make(map[string]func(*state, command) error),}
	cmds.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		fmt.Println("insufficient inputs. gator <command> <arguments>")
		os.Exit(1)
	}

	cmdName := args[1]
	cmdArguments := args[2:]

	cmd := command{
		Name: cmdName,
		Args: cmdArguments,
	}

	err = cmds.run(s, cmd)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
