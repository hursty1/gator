package main

import "fmt"

func (c *commands) run(s *state, cmd command) error {
	if handler, ok := c.cmd[cmd.name]; ok {
		return handler(s,cmd)
	}
	return fmt.Errorf("unknown command: %s", cmd.name)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmd[name] = f
}