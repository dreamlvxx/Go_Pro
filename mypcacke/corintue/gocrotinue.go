package main

import "fmt"

func main() {
	channelA := make(chan int)
	channelB := make(chan int)

	go func() {
		for x := 0; x < 10; x++ {
			channelA <- x
		}
	}()

	go func() {
		for b := 0; b < 20; b++ {
			channelB <- b
		}
	}()

	for true {
		fmt.Println(<-channelB)
	}

	close(channelB)
	close(channelA)
}
