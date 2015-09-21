package main

import "fmt"

/* the following program adheres to the 'no shared channels' principle, 
   AND is more fair in its scheduling. */

func mux(n int, c chan<- int) {
	for {
		c <- n
	}
}

func main() {
	c := make(chan int)
	d := make(chan int)
	go mux(1, c)
	go mux(2, d)
	for i := 0; i < 1000; i++ {
		select {
			case x := <-c: fmt.Println(x)
			case x := <-d: fmt.Println(x)
		}
	}
}
