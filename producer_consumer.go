package main

import (
	"fmt"
	"math/rand"
)

func main() {
	mainBlock := make(chan int)
	chanEven := make(chan int)
	chanUneven := make(chan int)

	go producerEven(chanEven)
	go producerUneven(chanUneven)
	go consumer(chanEven, chanUneven)
	<-mainBlock
}

func producerEven(chanEven chan int) {
	for {
		x := rand.Int()
		if x%2 == 0 {
			chanEven <- x
		}
	}
}

func producerUneven(chanUneven chan int) {
	for {
		y := rand.Int()
		if y%2 == 1 {
			chanUneven <- y
		}
	}
}

func consumer(chanEven chan int, chanUneven chan int) {
	for {
		x := <-chanEven
		y := <-chanUneven
		fmt.Println(x)
		fmt.Println(y)
	}
}
