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
		fmt.Println("Square is Alive!")
	} else {
		fmt.Println("Square is Dead!")
	}
}

func main() {
	d := Dead
	d.display()
}

