package main

import (
	"fmt"
	"math"
)

type searchLimits struct {
	depth    int
	nodes    uint64
	moveTime int // in milliseconds
	infinite bool
}

var limits searchLimits

func (s *searchLimits) init() {
	s.depth = 9999
	s.nodes = math.MaxUint64
	s.moveTime = 99999999999
	s.infinite = false
}

func (s *searchLimits) setDepth(d int) {
	s.depth = d
}

func (s *searchLimits) setMoveTime(m int) {
	s.moveTime = m
}

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
			case "go":
				tell("info string Im thinking")
				// TODO start the thinking process in the engine from "go"
			}
		}
	}()

	return frEng, toEng
}
