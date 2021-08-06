package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* Square stuff */

type square bool

// a square can be alive or dead
const (
	Alive square = true
	Dead  square = false
)

// convert to string
func (s square) String() string {
	if s == Alive {
		return "X"
	} else {
		return " "
	}
}

/* Board stuff */

type board struct {
	height      int
	width       int
	squares     [][]square
	nextSquares [][]square
}

func (b *board) init(width, height int) {
	b.height = height
	b.width = width

	b.squares = make([][]square, height)
	for i := 0; i < height; i++ {
		b.squares[i] = make([]square, width)
	}

	b.nextSquares = make([][]square, height)
	for i := 0; i < height; i++ {
		b.nextSquares[i] = make([]square, width)
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
			fmt.Print(b.squares[y][x])
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

// seperate buffer generated from current state:
func (b *board) setNext(x, y int, value square) {
	b.nextSquares[y][x] = value
}

func (b board) get(x, y int) square {
	// wrap board like a torus
	if y < 0 {
		y = b.height + y
	}
	if x < 0 {
		x = b.width + x
	}
	if y >= b.height {
		y = b.height - y
	}
	if x >= b.width {
		x = b.width - x
	}
	return b.squares[y][x]
}

// Apply the rules to one square
func (b board) getNumberOfNeighbors(x, y int) int {
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
func (b board) applyRules(x, y int) {
	n := b.getNumberOfNeighbors(x, y)

	// Any live cell
	if b.get(x, y) {
		if n < 2 { // with fewer than two live neighbours
			// dies, as if by underpopulation.
			b.setNext(x, y, Dead)
		} else if n > 3 { // with more than three live neighbours
			// dies, as if by overpopulation.
			b.setNext(x, y, Dead)
		} else { // with two or three live neighbours
			// lives on to the next generation.
			b.setNext(x, y, Alive)
		}
	} else {
		// Any dead cell with exactly three live neighbours
		// becomes a live cell, as if by reproduction.
		if n == 3 {
			b.setNext(x, y, Alive)
		} else {
			b.setNext(x, y, Dead)
		}
	}
}

// Progress the board one frame
func (b board) tick() {
	// generate next frame
	for y, line := range b.squares {
		for x := range line {
			b.applyRules(x, y)
		}
	}

	// copy over from frame buffer
	for y, line := range b.nextSquares {
		for x := range line {
			b.squares[y][x] = b.nextSquares[y][x]
		}
	}
}

func (b board) randomize() {
	s := rand.NewSource(time.Now().UnixNano())
	for y, line := range b.squares {
		for x := range line {
			r := rand.New(s)
			if r.Intn(2) == 1 {
				b.set(x, y, Alive)
			}
		}
	}
}

func main() {
	b := new(board)
	b.init(80, 10)

	// Glider:
	b.setActive([][]int{
		[]int{5, 5},
		[]int{6, 6},
		[]int{6, 7},
		[]int{5, 7},
		[]int{4, 7},
	})
	// random start:
	b.randomize()

	b.show()

	for {
		t0 := time.Now()
		b.tick()
		t1 := time.Now()
		b.show()
		fmt.Println("frame generated in", t1.Sub(t0))
		time.Sleep(100 * time.Millisecond)
	}
}
