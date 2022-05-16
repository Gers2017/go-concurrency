package main

import (
	"testing"

	"golang.org/x/tour/tree"
)

func TestWalk(t *testing.T) {
	ch := make(chan int)
	t1 := tree.New(1)
	go Walk(t1, ch)

	for x := 1; x <= 10; x++ {
		v := <-ch
		if x != v {
			t.Errorf("Expected %d, Got %d\nt1:%v", x, v, t1.String())
		}
	}
}

func TestSameTree(t *testing.T) {
	t1 := tree.New(1)
	t2 := tree.New(1)
	if !Same(t1, t2) {
		t.Errorf("t1 should have the same structure as t2\nt1: %v\nt2: %v", t1.String(), t2.String())
	}
}
