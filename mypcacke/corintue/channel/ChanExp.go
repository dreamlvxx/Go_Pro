package main

import "fmt"

func main() {
	counterchannnel := make(chan int)
	squarchannel := make(chan int)

	go counter(counterchannnel)
	go squar(counterchannnel, squarchannel)
	printer(squarchannel)
}

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squar(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
