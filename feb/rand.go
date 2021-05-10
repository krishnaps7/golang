package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var count = 0
	for count < 10 {
		var num = rand.Intn(100)
		fmt.Println(num)
		count++
	}
}
