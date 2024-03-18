package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	nP12     = 12
	nP       = 6
	WHITE    = color(0)
	BLACK    = color(1)
	startpos = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - "
)

type boardStruct struct {
	sq        [64]int      // array of squares -> board
	wbBB      [2]bitBoard  // 2 bitboards for white and black pieces
	pieceBB   [nP]bitBoard // a bitboard for each piece (p, r, n, k, q, b)
	King      [2]int       // white and black king positions
	ep        int          // en passant
	castlings              // the castlings
	stm       color        // side to move (either white or black)
	count     [nP12]int    // count how many pieces of each type we have
}

type (
	castlings uint
	color     int // either WHITE or BLACK (0/1)
)

var board = boardStruct{}

// returns a bitboard will all the pieces
func (b *boardStruct) allBB() bitBoard {
	return b.wbBB[0] | b.wbBB[1]
}

// clear the board, flags, bitboards etc
func (b *boardStruct) clear() {
	b.stm = WHITE
	b.sq = [64]int{}

	for ix := A1; ix <= H8; ix++ {
		b.sq[ix] = empty
	}

	b.ep = 0
	b.castlings = 0

	for ix := 0; ix < nP12; ix++ {
		b.count[ix] = 0
	}

	// bitBoards
	b.wbBB[WHITE], b.wbBB[BLACK] = 0, 0
	for ix := 0; ix < nP; ix++ {
		b.pieceBB[ix] = 0
	}
}

func (b *boardStruct) newGame() {
	b.stm = WHITE
	b.clear()
	parseFEN(startpos)
}

// parse a FEN string and setup the position
func parseFEN(FEN string) {
	fenIx := 0

	for row := 7; row >= 0; row-- {
		for sq := row * 8; sq < row*8-8; {

			char := string(FEN[fenIx])
			fenIx++
			if char == "/" {
				continue
			}

			if i, err := strconv.Atoi(char); err == nil {
				fmt.Println(i, "empty for sq", sq)
				sq += i
				continue
			}

			fmt.Println(char, "on sq", sq)
			sq++
		}
	}
}

// parse and make the moves in position command from GUI
func parseMvs(mvstr string) {
	mvs := strings.Split(mvstr, " ")

	for _, mv := range mvs {
		fmt.Println("make move", mv)
	}
}

// 6 piece types - no color (P)
const (
	Pawn int = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

// 12 pieces with color (P12)
const (
	wP = iota
	bP
	wN
	bN
	wB
	bB
	wR
	bR
	wQ
	bQ
	wK
	bK
	empty = 15
)

// square names
const (
	A1 = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1

	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2

	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3

	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4

	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5

	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6

	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7

	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)
