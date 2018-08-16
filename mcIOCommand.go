package main

import (
	"io/ioutil"
	"log"
	"os"
)

type mcIOCommand struct {
	Command
}

func (c *mcIOCommand) Prepare(parameter string) {
}

func (c *mcIOCommand) Payload() string {
	b, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Panic(err)
	}

	return string(b)
}

func (c *mcIOCommand) Receive(b []byte) (string, error) {
	return string(b), nil
}
