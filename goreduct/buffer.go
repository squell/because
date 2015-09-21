package main

import "fmt"

func buffer(n int, in chan int) chan int {
	if n == 0 {
		return in
	}
	tmp := make(chan int)
	go func() {
		for x := range in {
			tmp <- x
		}
	}()
	return buffer(n-1, tmp)
}

func main() {
	c := make(chan int)
	d := buffer(2, c)
	c <- 1
	c <- 1
	c <- 1
	fmt.Println(<-d)
	fmt.Println(<-d)
	fmt.Println(<-d)
}
