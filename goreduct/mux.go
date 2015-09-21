package main

import "fmt"

/* the following program doesn't schedule fairly between 1 and 2 if GOMAXPROCS=1; 
   mux(1,c) gets 66% of the cpu time
   it does schedule about fairly with a buffered channel (len=1);
   with buffered channel (len=2), mux(1,m) gets 100% of the cpu time. 

   with GOMAXPROCS>1, it is different again. */

func mux(n int, c chan<- int) {
	for {
		c <- n
	}
}

func main() {
	c := make(chan int,10)
	go mux(1, c)
	go mux(2, c)
	for i := 0; i < 1000; i++ {
		fmt.Println(<-c)
	}
}
