package main

import (
	"fmt"
	"time"
)

func rt1(c chan string) {
	c <- "gif"
	c <- "pic"
	c <- "png"
	c <- "zip"
	// fmt.Print("length of the channel is: ", len(c))
	close(c)
}

func main() {
	c := make(chan string, 5)
	d := make(chan string, 5)
	go rt1(c)
	time.Sleep(5 * time.Second)
	// fmt.Print("length of the channel is: ", len(c))
	for x := range c {
		fmt.Println(x)
	}
	d <- "krishna"
	d <- "renuka"
	d <- "radhika"
	d <- "laxmi"
	fmt.Println("Length of the d is ", len(d))
}
