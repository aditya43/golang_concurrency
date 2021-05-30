package main

import "fmt"

// Implement relaying of message with Channel Direction

func genMsg(ch1 chan<- string) {
	// send message on ch1
	ch1 <- "Hello Aditya"
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	// recv message on ch1
	recv := <-ch1

	// send it on ch2
	ch2 <- recv
}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// spin goroutines genMsg and relayMsg
	go genMsg(ch1)
	go relayMsg(ch1, ch2)

	// recv message on ch2
	recv := <-ch2
	fmt.Println(recv)
}
