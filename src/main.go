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
	if y < 0 {
		y = b.height + y
	}
	if x < 0 {
		x = b.width + x
	}
	return b.squares[y][x]
}

func (b *board) getRef(x, y int) *square {
	if y < 0 {
		y = b.height + y
	}
	if x < 0 {
		x = b.width + x
	}
	return &b.squares[y][x]
}

// Apply the rules to one square
func (b *board) getNumberOfNeighbors(x, y int) int {
	n := 0
	// dx and dy are deltas from the given (x, y) point,
	// Here is a graph:
	/*
		(x-1, y-1) | (x, y-1) | (x+1, y-1)
		(x-1, y  ) | (x, y  ) | (x+1, y  )
		(x-1, y+1) | (x, y+1) | (x+1, y+1)
	*/
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			// make sure we don't count the square we're checking around
			if x+dx != x || y+dy != y {
				// count alive squares
				if b.get(x+dx, y+dy) {
					n++
				}
			}
		}
	}
	return n
}

// Apply the rules to one square
func (b *board) applyRules(x, y int) {
	fmt.Println("Rules: ", x, y)
}

// Progress the board one frame
func (b board) tick() {
	for y, line := range b.squares {
		for x := range line {
			b.applyRules(x, y)
		}
	}
}

func main() {
	b := new(board)
	b.init(10, 10)
	b.setActive([][]int{
		// around (5,5):
		[]int{4, 4},
		[]int{5, 4},
		[]int{4, 5},
		[]int{6, 5},
		[]int{6, 6},
		[]int{5, 6},
		[]int{6, 4},
		[]int{4, 6},
		// around (0,0):
		[]int{0, 9},
		[]int{0, 1},
		[]int{1, 0},
		[]int{1, 1},
		[]int{1, 9},
		[]int{9, 0},
		[]int{9, 1},
		[]int{9, 9},
	})
	b.show()
	fmt.Println("5, 5 has", b.getNumberOfNeighbors(5, 5), "neighbors")
	fmt.Println("0, 0 has", b.getNumberOfNeighbors(0, 0), "neighbors")
}
