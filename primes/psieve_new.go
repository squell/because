// Prime number sieve. Concurrent version

package main

import "fmt"

const MAX_PRIME = 25000

func sieve(src <-chan int, out chan<- int) {
	p := <-src
	out <- p
	filt := make(chan int)
	go sieve(filt, out)
	for {
		n := <-src
		if n%p != 0 {
			filt <- n
		}
	}
}

func generate(init int, out chan<- int) {
	x := init
	for { 
		out <- x
		x++
	}
}

func main() {
	outp := make(chan int)
	inp := make(chan int)
	go sieve(inp, outp)
	go generate(2,inp)
	for {
		n := <-outp
		if n >= MAX_PRIME {
			break
		}
		fmt.Println(n)
	}
}
