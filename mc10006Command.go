package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

type mc10006Command struct {
	mcBaseCommand
	Modelno string
	Side    string
}

func (c *mc10006Command) Prepare(parameter string) {
	var modelno string
	var side string

	fmt.Sscanf(parameter, "%s%s", &modelno, &side)

	c.Modelno = modelno
	c.Side = side
}

func (c *mc10006Command) Payload() string {
	b, err := ioutil.ReadFile(filepath.Join(".", "data", "10006_request.xml"))

	if err != nil {
		log.Panic(err)
	}

	message := new(Message)

	err = xml.Unmarshal(b, message)

	if err != nil {
		log.Panic(err)
	}

	message.Body.ModelInfo.ModelNo = c.Modelno
	message.Body.ModelInfo.Side = c.Side

	m, err := xml.Marshal(message)

	if err != nil {
		log.Panic(err)
	}

	return string(m)
}

func (c *mc10006Command) Receive(b []byte) (string, error) {
	msg := new(Message)

	err := xml.Unmarshal(b, msg)

	if err != nil {
		return "", err
	}

	if "0" == msg.Body.Result.ErrorCode {
		msg580 := new(Message)

		msg580.Header = msg.Header
		msg580.Header.MessageClass = "580"

		msg580.Body.BOM = msg.Body.BOM
		msg580.Body.Resource = msg.Body.Resource
		msg580.Body.ModelInfo = msg.Body.ModelInfo

		m, err := xml.Marshal(msg580)

		if err != nil {
			return "", err
		}

		return string(m), nil
	}

	return "", fmt.Errorf("ERROR : %s", msg.Body.Result.ErrorText)
}
