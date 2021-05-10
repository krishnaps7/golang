package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 1000
	for sum < 1001 {
		fmt.Println(sum)
		sum -= 1
		time.Sleep(time.Second)
		fmt.Println(time.Now())
		if sum < 995 {
			break
		}
	}
	fmt.Println(time.Now())
}
