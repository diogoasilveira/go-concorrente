package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxCapacity = 10

func create_req(req chan int) {
	for {
		time.Sleep(1 * time.Second)
		valor := rand.Intn(100)
		req <- valor
	}
}

func exec_req(req chan int) {
	for range req {
		v1 := <-req
		fmt.Println(v1)
	}
}

func main() {
	block := make(chan int)
	requests := make(chan int, maxCapacity)
	go create_req(requests)
	for i := 0; i < maxCapacity; i++ {
		go exec_req(requests)
	}

	<-block
}
