package main

import (
	"fmt"
)

type square bool

// a square can be alive or dead
const (
	Alive square = true
	Dead square = false
)

func (s square) display() {
	if s == Alive {
		fmt.Print("X")
	} else {
		fmt.Print(" ")
	}
}

type board struct {
	height int
	width int
	squares [][]square
}

func (b *board) init(height, width int) {
	b.height = height
	b.width = width

	b.squares = make([][]square, height)
	for i := 0; i < height; i++ {
		    b.squares[i] = make([]square, width)
	}
}

func (b board) printHeader() {
	fmt.Print("+")
	for x := 0; x < b.width; x++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}

func (b board) display() {
	b.printHeader()
	for y := 0; y < b.height; y++ {
		fmt.Print("|")
		for x := 0; x < b.width; x++ {
			b.squares[y][x].display()
		}
		fmt.Println("|")
	}
	b.printHeader()
}


func main() {
	b := new(board)
	b.init(10,10)
	b.display()
}

