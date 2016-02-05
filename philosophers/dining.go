package main

import "fmt"

func Lefty(id int, eating chan int, left chan int, right chan int) {
	for {
		y := <-left
		x := <-right
		eating <- id
		right <- x
		left <- y
	}
}

func Philosopher(id int, eating chan int, left chan int, right chan int) {
	for {
		x := <-right
		y := <-left
		eating <- id
		right <- x
		left <- y
	}
}

func Table(id int, fork chan int) {
	for {
		fork <- id
		<-fork
	}
}

func main() {
	var forks [5]chan int
	for i := range forks {
		forks[i] = make(chan int)
		go Table(i, forks[i])
	}

	var eating = make(chan int)
	for i := range forks {
		if i == 0 {
			go Lefty(i, eating, forks[i], forks[(i+1)%5])
		} else {
			go Philosopher(i, eating, forks[i], forks[(i+1)%5])
		}
	}

	for {
		fmt.Printf("philosopher #%d eating\n", <-eating)
	}
}
