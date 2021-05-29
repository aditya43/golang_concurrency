package main

import "fmt"

func main() {
	ch := make(chan int)

	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)

	res := <-ch
	// TODO: get the value computed from goroutine
	fmt.Printf("computed value %v\n", res)
}

// Original Code:
// package main

// func main() {
// 	go func(a, b int) {
// 		c := a + b
// 	}(1, 2)
// 	// TODO: get the value computed from goroutine
// 	// fmt.Printf("computed value %v\n", c)
// }
