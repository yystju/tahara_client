package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"sync"
)

// Command is the interface to generate the payload.
type Command interface {
	Prepare(parameter string)
	Payload() string
	Receive(msg []byte) (string, error)
}

func doReceiveAndClose(conn *net.TCPConn, c Command, latch chan int) {
	scanner := bufio.NewScanner(conn)

	buff := make([]byte, 8*1024*1024)

	scanner.Buffer(buff, len(buff))

	scanner.Split(packageSpliter)

	for scanner.Scan() {
		if scanner.Err() != nil {
			log.Printf("error(%v) break...", scanner.Err())
			break
		}

		b := scanner.Bytes()

		if len(b) > 0 {
			log.Println("-received-")

			i := bytes.IndexByte(b, '\x01')

			ret, err := c.Receive(b[i+1:])

			if err != nil {
				log.Panic(err)
			}

			fmt.Println(ret)

			latch <- 1
		}
	}

	log.Println("RECEIVER DONE.")
}

func doPost(conn *net.TCPConn, c Command, latch chan int) {
	s := c.Payload()

	for i := 0; i < count; i++ {
		conn.Write([]byte(fmt.Sprintf("\x02%d\x01%s\x03", len(s), s)))

		<-latch
	}

	log.Println("POST DONE.")
}

func doCommand(addr string, c Command, w *sync.WaitGroup) {
	raddr, err := net.ResolveTCPAddr("tcp", addr)

	if err != nil {
		log.Panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, raddr)

	if err != nil {
		log.Panic(err)
	}

	log.Println("CONNECTION: ", conn.RemoteAddr(), conn.LocalAddr())

	latch := make(chan int)
	go doReceiveAndClose(conn, c, latch)
	doPost(conn, c, latch)

	conn.Close()

	w.Done()

	log.Println("-COMMAND DONE.-")
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
