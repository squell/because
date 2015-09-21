package main

import "fmt"

func mux(n int, in <-chan int, out chan<- int) {
	for {
		out <- n+ <- in
	}
}

// in theory, this program doesn't always terminate

func main() {
	c, d := make(chan int), make(chan int)
	go mux(10,c,d)
	go mux(100,d,c)
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Println(<-d)
	}
}

