package main

import (
	"fmt"
	"time"
)

func main() {
	tic := time.Tick(100 * time.Millisecond)
	tac := time.Tick(150 * time.Millisecond)
	boom := time.After(400 * time.Millisecond)

	for {
		select {
		case <-tic:
			fmt.Println("tic")
		case <-tac:
			fmt.Println("tac")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default: // When none of the above channels are ready
			fmt.Println("...")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
