package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	addr        string
	command     string
	parameter   string
	count       int
	concurrency int

	commands = make(map[string]Command)

	done = make(chan int)
)

func init() {
	mc10006 := new(mc10006Command)
	commands["10006"] = mc10006

	keys := make([]string, 0, len(commands))
	for k := range commands {
		keys = append(keys, k)
	}

	flag.StringVar(&addr, "s", ":58060", "The server address \"host:port\".")
	flag.StringVar(&command, "c", "-", fmt.Sprintf("The command in %v. \"-\" indicates not a command mode.", keys))
	flag.StringVar(&parameter, "p", "", "The parameters (space-separated) for the command.")
	flag.IntVar(&count, "t", 1, "Number of times to loop.")
	flag.IntVar(&concurrency, "n", 1, "Number of threads.")
	flag.Parse()
}

func main() {
	f, err := os.OpenFile(filepath.Join(".", "client.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	log.SetOutput(f)

	w := new(sync.WaitGroup)

	w.Add(concurrency)

	if command == "-" {
		cmd := new(mcIOCommand)
		cmd.Prepare("")
		for i := 0; i < concurrency; i++ {
			go doCommand(addr, cmd, w)
		}
	} else { //Command mode
		commands[command].Prepare(parameter)
		for i := 0; i < concurrency; i++ {
			doCommand(addr, commands[command], w)
		}
	}

	w.Wait()

	log.Println("MAIN DONE.")
}
