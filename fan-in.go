package main

func request_stream() chan string {
	channel := make(chan string)
	return channel
}

func ingest(in chan string) {

}

func main() {
	ch1 := request_stream()
	ch2 := request_stream()
	ch3 := make(chan string)

	go ingest(ch3)

	for {
		select {
		case v1 := <-ch1:
			ch3 <- v1
		case v2 := <-ch2:
			ch3 <- v2
		}
	}
}
