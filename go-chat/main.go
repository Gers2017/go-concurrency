package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getMessage(n int, ch chan string) { // <- sender
	names := []string{"Joe", "Sam", "Steve", "Alice", "Bob", "Eve"}
	for i := 0; i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		ri := rand.Intn(len(names))
		name := names[ri]
		ch <- "Hello " + name
	}
	close(ch) // Closes the channel
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan string, 16)

	go getMessage(cap(ch), ch)
	for msg := range ch { // <- receiver
		fmt.Println(msg)
	}
}
