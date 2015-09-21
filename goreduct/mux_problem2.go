package main

import "fmt"

func mux(n int, in <-chan int, out chan<- int) {
	for {
		out <- n+ <- in
	}
}

// rewritten to conform to the "one writer/reader" property; however,
// this program can (and will) produce deadlocks; so this is not an
// equivalent solution

func main() {
	c, d := make(chan int), make(chan int)
        e := make(chan int)
	go mux(10,c,d)
	go mux(100,d,e)
	for i := 0; i < 10; i++ {
		c <- i
		select {
			case x:=<-e: fmt.Println(x)
			case x:=<-e: c<-x // will deadlock
		}
	}
}

