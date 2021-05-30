package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for i := 1; i < 4; i++ {
			time.Sleep(1 * time.Second)
			ch <- fmt.Sprintf("Message %d", i)
		}

	}()

	// if there is no value on channel, do not block.
	for i := 0; i < 10; i++ {
		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			fmt.Println("No message received!")
		}

		// Do some processing..
		fmt.Println("processing..")
		time.Sleep(1500 * time.Millisecond)
	}
}
