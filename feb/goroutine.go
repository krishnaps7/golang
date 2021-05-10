package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println("Welcome to concurrency,", str)
		time.Sleep(time.Second)
	}
	// return str
}

func main() {
	go display("krishna")
	display("Sabaris")
	go func(name string) {
		time.Sleep(time.Second)
		fmt.Println("your name is:", name)
	}("Krishna Sabaris PS")

	time.Sleep(time.Second)
}
