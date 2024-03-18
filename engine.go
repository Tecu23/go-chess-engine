package main

import "fmt"

// engine is a go routine with input and output channels
func engine() (frEng, toEng chan string) {
	fmt.Println("info string Hello from engine")

	frEng = make(chan string)
	toEng = make(chan string)

	go func() {
		for cmd := range toEng {
			tell("info string engine got ", cmd)
			switch cmd {
			case "stop":
			case "quit":
			}
		}
	}()

	return frEng, toEng
}
