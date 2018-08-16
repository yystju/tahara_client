package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
)

// Command is the interface to generate the payload.
type Command interface {
	Prepare(parameter string)
	Payload() string
	Receive(msg []byte) (string, error)
}

func doReceiveAndClose(conn *net.TCPConn, c Command) {
	scanner := bufio.NewScanner(conn)

	buff := make([]byte, 8*1024*1024)

	scanner.Buffer(buff, len(buff))

	scanner.Split(packageSpliter)

	for scanner.Scan() {
		b := scanner.Bytes()
		if len(b) > 0 {
			i := bytes.IndexByte(b, '\x01')

			ret, err := c.Receive(b[i+1:])

			if err != nil {
				log.Panic(err)
			}

			fmt.Println(ret)

			conn.Close()

			done <- 1
		}
	}
}

func doPost(conn *net.TCPConn, c Command) {
	s := c.Payload()
	conn.Write([]byte(fmt.Sprintf("\x02%d\x01%s\x03", len(s), s)))
}

func doCommand(addr string, c Command) {
	raddr, err := net.ResolveTCPAddr("tcp", addr)

	if err != nil {
		log.Panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, raddr)

	if err != nil {
		log.Panic(err)
	}

	log.Println("CONNECTION: ", conn.RemoteAddr(), conn.LocalAddr())

	go doReceiveAndClose(conn, c)
	doPost(conn, c)
}

func packageSpliter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == '\x03' {
			return i + 1, data[:i], nil
		}
	}

	if atEOF {
		return 0, data, bufio.ErrFinalToken
	}

	return 0, nil, nil
}
