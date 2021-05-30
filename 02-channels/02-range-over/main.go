package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		defer close(ch) // Close channel
		for i := 0; i < 6; i++ {
			// send iterator over channel
			ch <- i
		}
	}()

	// range over channel to recv values
	for v := range ch {
		fmt.Println(v)
	}
}
