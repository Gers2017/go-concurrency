package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {
	ch := make(chan int)
	t := tree.New(1)
	Traverse(t)
	fmt.Println()
	go Walk(t, ch)

	for x := 0; x < 10; x++ {
		fmt.Printf("%d ", <-ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Traverse(t *tree.Tree) {
	if t == nil {
		return
	}
	Traverse(t.Left)
	fmt.Print("->", t.Value)
	Traverse(t.Right)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	Ach := make(chan int)
	Bch := make(chan int)

	go Walk(t1, Ach)
	go Walk(t2, Bch)

	for x := 0; x < 10; x++ {
		a := <-Ach
		b := <-Bch
		if a != b {
			return false
		}
	}

	return true
}
