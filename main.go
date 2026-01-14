package main

import (
	"database/sql"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}
	defer db.Close()

	dbQueries := database.New(db)

	s := &state{
		config:   &cfg,
		database: dbQueries,
	}

	cmds := commands{registeredCommands: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerFeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", handlerFollow)

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
