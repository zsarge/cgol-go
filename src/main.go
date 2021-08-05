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

func (s *square) invert() {
	if *s == Alive {
		*s = Dead
	} else {
		*s = Alive
	}
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

func (b board) setActive(arr [][]int) {
	for y := range arr {
		b.set(arr[y][0], arr[y][1], Alive)
	}
}

func (b *board) set(x, y int, value square) {
	b.squares[y][x] = value
}

func (b *board) get(x, y int) square {
	return b.squares[y][x]
}

func (b *board) getRef(x, y int) *square {
	return &b.squares[y][x]
}

// Progress the board one frame
func (b *board) tick() {
	for y, line := range b.squares {
		for x := range line {
			b.squares[y][x].invert()
		}
	}
}

func main() {
	b := new(board)
	b.init(10, 10)
	b.setActive([][]int{
		[]int{4, 4},
		[]int{5, 4},
		[]int{4, 5},
		[]int{5, 5},
	})
	b.show()
	b.getRef(5, 5).invert()
	fmt.Println("invert")
	b.show()
	fmt.Println("tick")
	b.tick()
	b.show()
}
