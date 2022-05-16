package main

import (
	"fmt"
	"math/rand"
	"time"
)

func PingRandomInt(ch, stopch chan int) {
	for {
		r := rand.Int()
		time.Sleep(time.Millisecond * 100)
		select {
		case ch <- r:
			rand_number := rand.Int()
			r = int(rand_number * rand_number)
		case <-stopch:
			fmt.Println("Stop!")
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)
	stopch := make(chan int)

	go func() {
		for i := 0; i < 12; i++ {
			fmt.Println(<-ch)
		}
		stopch <- 0
	}()

	PingRandomInt(ch, stopch)
}
