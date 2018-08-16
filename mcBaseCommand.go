package main

import (
	"fmt"
)

type mcBaseCommand struct {
	Command
}

func (c *mcBaseCommand) Prepare(parameter string) {
}

func (c *mcBaseCommand) Payload() string {
	return ""
}

func (c *mcBaseCommand) Receive(b []byte) (string, error) {
	return "", fmt.Errorf("ERROR : %s", "Not implemented.")
}
