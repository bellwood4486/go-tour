package main

import "fmt"

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n+6; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)
	for i := range ch {
		fmt.Printf("%v (len: %v, cap: %v)\n", i, len(ch), cap(ch))
	}
}