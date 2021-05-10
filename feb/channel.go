package main

import (
	"fmt"
	"strconv"
	"time"
)

func myfun(mychan chan string) {
	for v := 0; v < 4; v++ {
		mychan <- "Go for Geek"
		time.Sleep(1 * time.Second)
	}
	close(mychan)
}
func test(c chan string) {
	for i := 0; i < 3; i++ {
		c <- "some dump" + strconv.Itoa(i)

	}
}

func main() {
	ch := make(chan string)
	c := make(chan string)
	go test(c)
	fmt.Println("data in c", <-c)
	go myfun(ch)
	for {
		data, ok := <-ch
		if ok == false {
			fmt.Println("Channel is Closed")
			break
		}
		fmt.Println("Channel Open:", data, ok)
	}
}
