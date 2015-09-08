// Prime number sieve. Concurrent version

package main

import "fmt"

const MAX_PRIME = 25000

func sieve(in <-chan int, out chan<- int) {
    myPrime := <-in
    out <- myPrime
    filt := make(chan int)
    go sieve(filt, out)
    for n := range in {
        if(n % myPrime != 0) {
            filt <- n
        }
    }
}

func main() {
    outp := make(chan int)
    inp := make(chan int)
    go sieve(inp, outp)
    go func() {
        for x:=2; ; x++ { inp <- x }
    }()
    for n := range outp {
        if(n>=MAX_PRIME) { break }
        fmt.Println(n)
    }
}
