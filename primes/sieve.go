// Prime number sieve. Single threaded.

package main

import "fmt"

const MAX_PRIME = 25000
var primes [MAX_PRIME]int

func main() {
    end := 0
outer:
    for n:=2; n<MAX_PRIME; n++ {
        for i:=0; i < end; i++ {
            if(n % primes[i] == 0) { continue outer }
        }
        primes[end] = n
        end++
        fmt.Printf("%d\n", n);
    }
}
