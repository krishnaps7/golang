package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	go sleepyGoats()
// 	fmt.Println("I'm in main")
// 	time.Sleep(4 * time.Second)
// 	fmt.Println("in main again")
// }

// func sleepyGoats() {
// 	fmt.Println("I am in goat")
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("...Snore...")
// }

//func main() {
//	c := make(chan int)
//	fmt.Println("At main gate")
//	for i := 0; i < 5; i++ {
//		go sleepyGoat(i, c)
//	}
//	for i := 0; i < 5; i++ {
//		gopherId := <-c
//		fmt.Println("gopher", gopherId, "has finished sleeping")
//	}
//	time.Sleep(4 * time.Second)
//	fmt.Println("At exit gate")
//}

//func sleepyGoat(n int, c chan int) {
//	//fmt.Println("At sleep mode - %v", n)
//	time.Sleep(3 * time.Second)
//	fmt.Println("..snore.. - %v", n)
//	c <- n
//}
func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGoat(i, c)
	}
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopherId := <-c:
			fmt.Println("gopher", gopherId, "has finied sleeping")
		case <-timeout:
			fmt.Println("My patience ran out")
			return
		}
	}
}

func sleepyGoat(n int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- n
}
