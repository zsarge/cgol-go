package main

import (
	"fmt"
)

/* Square stuff */

type square bool

// a square can be alive or dead
const (
	Alive square = true
	Dead  square = false
)

func (s square) show() {
	fmt.Print(s)
}

func (s square) String() string {
	if s == Alive {
		return "X"
	} else {
		return " "
	}
}

/* Board stuff */

type board struct {
	height  int
	width   int
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

func (b board) show() {
	b.printHeader()
	for y := 0; y < b.height; y++ {
		fmt.Print("|")
		for x := 0; x < b.width; x++ {
			b.squares[y][x].show()
		}
		fmt.Println("|")
	}
	b.printHeader()
}

func (b *board) set(x, y int, value square) {
	b.squares[y][x] = value
}

func (b *board) get(x, y int) square {
	return b.squares[y][x]
}

func main() {
	b := new(board)
	b.init(10, 10)
	b.set(4, 4, Alive)
	b.set(4, 5, Alive)
	b.set(5, 4, Alive)
	b.set(5, 5, Alive)
	fmt.Println(b.get(5, 5))
	b.show()
}
