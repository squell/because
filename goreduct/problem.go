package main

import "fmt"

// the problem with channels part 1

// the following function might be able to read its own messages, deadlocking main,
// if c is buffered (and that information isn't part of the type!), but if
// c is unbuffered, everything will be fine. this is a counterexample to the claim that
// if a program behaves correctly with unbuffered channels, it will behave correctly
// with buffered channels.

// also note that the last line of sub() isn't guaranteed to be executed
// (it will not be if GOMAXPROCS=1, it will be if GOMAXPROCS>1)

// to eliminate these cases, we should disallow "input-output-aliassing";
// but bidirectional channels are always input-output aliassed; so, we should
// statically *PROVE* that aliassing cannot occur and that a channel is always 
// only used in 'one direction'

func sub(c chan int) {
	c <- 5
	x := <-c
	fmt.Println("sub:", x)
}

func main() {
	c := make(chan int,42)
	go sub(c)
	fmt.Println("main:", <-c)
	c<-6
	for { 
	}
}

