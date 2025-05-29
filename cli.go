package main

import (
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) !=1 {
			fmt.Println("username is required")
			os.Exit(1)
		}
	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to %s\n", cmd.args[0])
	return nil
}