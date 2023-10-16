package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	mainBlock := make(chan int)
	closed := make(chan int)
	n := 3
	generateRoutines(n, closed)
	z, ok := <-closed
	z++
	if !ok {
		println(n)
	}
	<-mainBlock
}

func routine() {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Println("mimir")
}

func generateRoutines(n int, c chan int) {
	for i := 0; i < n; i++ {
		go routine()
	}
	close(c)
}
