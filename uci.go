package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// tell is used to communicate between the go routines
var tell func(text ...string)

func uci(frGUI chan string, myTell func(text ...string)) {
	tell = myTell
	tell("info string Hello from uci")

	// to channels, from engine and to engine
	frEng, toEng := engine()

	quit := false // signal when to quit

	cmd := "" // command received from user

	bm := "" // best move received from engine

	for quit == false {

		select {
		case cmd = <-frGUI:
		case bm = <-frEng:
			handleBm(bm)
			continue
		}

		switch cmd {
		case "uci":
		case "stop":
			handleStop(toEng)
		case "quit", "q":
			quit = true
			continue

		}
	}
}

// Handle Best Move received from engine
func handleBm(bm string) {
	tell(bm)
}

// Handle Stop -> send to the engine the signal to stop
func handleStop(toEng chan string) {
	toEng <- "stop"
}

// input function that read from the STDIN and returns a string channel with the values read
func input() chan string {
	line := make(chan string)

	go func() {
		var reader *bufio.Reader
		reader = bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
			text = strings.TrimSpace(text)
			if err != io.EOF && len(text) > 0 {
				line <- text
			}
		}
	}()

	return line
}

func mainTell(text ...string) {
	toGUI := ""

	for _, t := range text {
		toGUI += t
	}

	fmt.Println(toGUI)
}
