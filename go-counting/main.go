package main

import "fmt"

func main() {

	ch := make(chan int)
	go count_part1(ch)
	go count_part2(ch)

	total_a, total_b := <-ch, <-ch

	fmt.Println("The result is:", total_a+total_b)

}

func count_part1(ch chan int) {
	a := 0
	for i := 1; i <= 500000; i++ {
		a += i
	}
	ch <- a
}

func count_part2(ch chan int) {
	b := 0
	for i := 500001; i <= 1000000; i++ {
		b += i
	}
	ch <- b
}
