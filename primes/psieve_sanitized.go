// Prime number sieve. Concurrent version

package main

import "fmt"

const MAX_PRIME = 25000

// original concurrent sieve() did not satisfy 'sanitized pipe' requirements

func prefix(head int, src <-chan int, dst chan<- int) {
	dst <- head
	for {
		n := <-src
		dst <- n
	}
}

func sieve(src <-chan int, dst chan<- int) {
	p := <-src
	filt := make(chan int)
	dst2 := make(chan int)
	go prefix(p, dst2, dst)
	go sieve(filt, dst2)
	for {
		n := <-src
		if n%p != 0 {
			filt <- n
		}
	}
}

func generator(out chan<- int) {
	for x := 2; ; x++ {
		out <- x
	}
}

func main() {
	outp := make(chan int)
	inp := make(chan int)
	go sieve(inp, outp)
	go generator(inp)
	for {
		n := <- outp
		if n >= MAX_PRIME {
			break
		}
		fmt.Println(n)
	}
}
