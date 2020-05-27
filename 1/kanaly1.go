package main

import (
	"fmt"
	"time"
)

func main() {
	commChannel, endChannel := outsideFunc()
	commChannel <- "testing comm one two three"
	fmt.Println("sent")
	commChannel <- "now a backflip"
	fmt.Println("sent again")
	endChannel <- struct{}{}
	fmt.Println("finito")
}

func mainloop(eventsChannel chan string, closingChannel chan struct{}) {
	fmt.Println("starting inner loop")
	for {
		select {
		case x := <-eventsChannel:

			fmt.Printf("to do \n %v \n", x)

		case end := <-closingChannel:

			fmt.Println("go down", end)

		}
	}
}

func outsideFunc() (chan string, chan struct{}) {
	eventsChannel := make(chan string)
	closingChannel := make(chan struct{})
	go mainloop(eventsChannel, closingChannel)
	go func() {
		for i := 0; i < 100000; i++ {
			fmt.Println("waiting for incomming message")
			time.Sleep(1 * time.Millisecond)
		}
	}()
	return eventsChannel, closingChannel
}
