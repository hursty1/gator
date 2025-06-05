package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/hursty1/gator/internal/database"

	"github.com/hursty1/gator/internal/config"
	_ "github.com/lib/pq"
)
func main() {

	c, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
		return
	}
	db, err := sql.Open("postgres", c.Db_url)
	if err != nil {
		fmt.Errorf("Error connecting")
		return
	}

	dbQueries := database.New(db)

	s := state{
		db: dbQueries,
		config: &c,

	}

	cli := commands{
		cmd: make(map[string]func(*state, command) error),
	}

	

	cli.register("login", handlerLogin)
	cli.register("config", showConfig)
	cli.register("register", register)
	cli.register("reset", reset)
	cli.register("users", middlewareLoggedIn(list))
	cli.register("agg", agg)
	cli.register("addfeed", middlewareLoggedIn(addfeed))
	cli.register("feeds", feeds)
	cli.register("follow", middlewareLoggedIn(follow))
	cli.register("following", middlewareLoggedIn(following))
	cli.register("unfollow", middlewareLoggedIn(unfollow))
	cli.register("browse", middlewareLoggedIn(browse))

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		os.Exit(1)
	}

	cli_cmd := args[1]
	cli_args := args[2:]
	
	err = cli.run(&s, command{name: cli_cmd, args: cli_args})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Program exit")
	os.Exit(0)

}