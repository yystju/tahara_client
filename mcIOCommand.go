package main

import (
	"io/ioutil"
	"log"
	"os"
)

type mcIOCommand struct {
	Command
	payload string
}

func (c *mcIOCommand) Prepare(parameter string) {
	if parameter == "" {
		b, err := ioutil.ReadAll(os.Stdin)

		if err != nil {
			log.Panic(err)
		}

		c.payload = string(b)
	} else {
		c.payload = parameter
	}
}

func (c *mcIOCommand) Payload() string {
	return c.payload
}

func (c *mcIOCommand) Receive(b []byte) (string, error) {
	return string(b), nil
}
