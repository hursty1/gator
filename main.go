package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hursty1/gator/internal/config"
)
func main() {

	c, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
		return
	}
	s := state{
		config: &c,
	}

	cli := commands{
		cmd: make(map[string]func(*state, command) error),
	}

	cli.register("login", handlerLogin)

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

}