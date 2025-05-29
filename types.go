package main

import "github.com/hursty1/gator/internal/config"

type state struct {
	config 	*config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmd map[string]func(*state, command) error
}