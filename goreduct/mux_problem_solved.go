package main

import "fmt"

func mux(n int, in <-chan int, out chan<- int) {
	for {
		out <- n+ <- in
	}
}

// rewritten to conform to the "one writer/reader" property;
// sort of equivalent.

func main() {
	c, d := make(chan int), make(chan int)
        e, f := make(chan int), make(chan int)
	go mux(10,c,d)
	go mux(100,e,c)
        go func(in <-chan int, out1 chan<- int, out2 chan<- int) {
		for {
			select {
				case x := <-in: out1 <- x
				case x := <-in: out2 <- x
			}
		}
	}(d,e,f)
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Println(<-f)
	}
}

