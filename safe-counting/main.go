package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := SafeCounter{value: 0}
	count_part1(&c)
	count_part2(&c)

	time.Sleep(time.Millisecond * 50)
	fmt.Println(c.Value())
}

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Inc(amount int) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.value
	c.value += amount
	c.mu.Unlock()
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.value
	defer c.mu.Unlock()
	// We can also use defer to ensure the mutex will be unlocked as in the Value method
	return c.value
}

func count_part1(c *SafeCounter) {
	a := 0
	for i := 1; i <= 500000; i++ {
		a += i
	}
	go c.Inc(a)
}

func count_part2(c *SafeCounter) {
	b := 0
	for i := 500001; i <= 1000000; i++ {
		b += i
	}
	go c.Inc(b)
}

func count_repeated(c *SafeCounter) {
	for i := 1; i <= 500000; i++ {
		go c.Inc(i)
	}

	for i := 500001; i <= 1000000; i++ {
		go c.Inc(i)
	}
}
