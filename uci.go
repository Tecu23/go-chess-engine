package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	trim = strings.TrimSpace
	low  = strings.ToLower
)

var saveBm = ""

func uci(inp chan string) {
	// to channels, from engine and to engine
	frEng, toEng := engine()

	bInfinite := false

	var cmd string // command received from user
	var bm string  // best move received from engine

	tell("info string Hello from uci")

	quit := false // signal when to quit
	for !quit {

		select {
		case cmd = <-inp:
			tell("info string uci got ", cmd)
		case bm = <-frEng:
			handleBm(bm, bInfinite)
			continue
		}

		words := strings.Split(cmd, " ")

		words[0] = trim(low(words[0]))

		switch words[0] {
		case "uci":
			handleUci()
		case "stop":
			handleStop(toEng, &bInfinite)
		case "quit", "q":
			handleQuit(toEng)
			quit = true
			continue

		}
	}
}

func handleUci() {
	tell("id name BinGo")
	tell("id author Carokanns")

	tell("option name Hash type spin default 128 min 16 max 1024")
	tell("option name Threads type spin default 1 min 1 max 16")
	tell("uciok")
}

func handleStop(toEng chan string, bInfinite *bool) {
	if *bInfinite {
		if saveBm != "" {
			tell(saveBm)
			saveBm = ""
		}

		toEng <- "stop"
		*bInfinite = false
	}
}

func handleQuit(toEng chan string) {
	toEng <- "stop"
}

// Handle Best Move received from engine
func handleBm(bm string, bInfinite bool) {
	if bInfinite {
		saveBm = bm
		return
	}
	tell(bm)
}

// input function that read from the STDIN and returns a string channel with the values read
func input() chan string {
	line := make(chan string)

	var reader *bufio.Reader
	reader = bufio.NewReader(os.Stdin)

	go func() {
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

func tell(text ...string) {
	toGUI := ""

	for _, t := range text {
		toGUI += t
	}

	fmt.Println(toGUI)
}
