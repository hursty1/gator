package main

import (
	"github.com/hursty1/gator/internal/config"
	"github.com/hursty1/gator/internal/database"
)

type state struct {
	db  *database.Queries
	config 	*config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmd map[string]func(*state, command) error
}