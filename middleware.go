package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hursty1/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		usr, err := s.db.GetUser(context.Background(), s.config.Current_user_name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
			return err
		}
		return handler(s, cmd, usr)
	}
}