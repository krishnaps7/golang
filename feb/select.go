package main

import (
	"fmt"
	"time"
)

func portal1(channel1 chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("portal1 woken up")
	channel1 <- "welcome to channel 1"
}

func portal2(channle2 chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("portal2 woken up")
	channle2 <- "welcome to channel 2"
}

func main() {
	R1 := make(chan string)
	R2 := make(chan string)

	go portal1(R1)
	go portal2(R2)

	select {
	case op1 := <-R1:
		fmt.Println(op1)
	case op2 := <-R2:
		fmt.Println(op2)
	}
	time.Sleep(4 * time.Second)
}
