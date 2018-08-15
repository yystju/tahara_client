package main

import (
	"encoding/xml"
	"fmt"
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

func (c *mcIOCommand) Receive(msg *Message) (string, error) {
	if "0" == msg.Body.Result.ErrorCode {
		m, err := xml.Marshal(msg)

		if err != nil {
			log.Panic(err)
		}

		return string(m), nil
	}

	return "", fmt.Errorf("ERROR : %s", msg.Body.Result.ErrorText)
}
