package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// what is the output
	//TODO: fix the issue.

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		/*
			// Original Code Without Fix: Output: 4, 4, 4
				go func() {
					defer wg.Done()
					fmt.Println(i)
				}()
		*/
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
