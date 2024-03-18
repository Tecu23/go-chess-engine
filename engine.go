package main

// engine is a go routine with input and output channels
func engine() (frEng, toEng chan string) {
	tell("Hello from engine")

	frEng = make(chan string)
	toEng = make(chan string)

	go func() {
		for cmd := range toEng {
			switch cmd {
			case "stop":
			case "quit":
			}
		}
	}()

	return frEng, toEng
}
