package main

import "fmt"

func Lefty(id int, eating chan<- int, left <-chan int, right <-chan int) {
	for eat := true; ; eat = !eat {
		<-left
		<-right
		if eat {
			eating <- id
		}
	}
}

func Philosopher(id int, eating chan int, left <-chan int, right <-chan int) {
	for eat := true; ; eat = !eat {
		<-right
		<-left
		if eat {
			eating <- id
		}
	}
}

func Table(id int, left chan<- int, right chan<- int) {
	for {
		select {
		case left <- id:
			left <- id
		case right <- id:
			right <- id
		}
	}
}

func main() {
	var forks [2][5]chan int
	for i := 0; i < 5; i++ {
		forks[0][i] = make(chan int)
		forks[1][i] = make(chan int)
		go Table(i, forks[0][i], forks[1][i])
	}

	var eating = make(chan int)
	for i := 0; i < 5; i++ {
		if i == 0 {
			go Lefty(i, eating, forks[0][i], forks[1][(i+1)%5])
		} else {
			go Philosopher(i, eating, forks[0][i], forks[1][(i+1)%5])
		}
	}

	for {
		fmt.Printf("philosopher #%d eating\n", <-eating)
	}
}
