package main

import "fmt"

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, val := range nums {
			out <- val
		}
	}()

	return out
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for val := range in {
			out <- val * val
		}
	}()

	return out
}

func main() {
	// set up the pipeline

	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.

	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// Method #1 to setup a pipeline
	ch := generator(2, 3)
	out := square(ch)

	for val := range out {
		fmt.Println(val)
	}
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// Another way to set up pipeline:
	for val := range square(generator(2, 3)) {
		fmt.Println(val)
	}
}
