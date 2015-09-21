package main

import "fmt"

func mux(n int, in <-chan int, out chan<- int) {
	for {
		fmt.Println("mux:",n)
		out <- <- in
	}
}

/* original program:
func main() {
	c, d := make(chan int), make(chan int)
	go mux(1,c,d)
	go mux(2,c,d)
	for i := 0; i < 10; i++ {
		c <- i
		fmt.Println(<-d)
	}
}
*/

func main() {
	c0, d0 := make(chan int), make(chan int)
	c1, d1 := make(chan int), make(chan int)
	go mux(1,c0,d0)
	go mux(2,c1,d1)
	for i := 0; i < 10; i++ {
		select {
			case c0 <- i: 
			case c1 <- i: 
		}
		select {
			case x:=<-d0: fmt.Println(x)
			case x:=<-d1: fmt.Println(x)
		}
	}
}
