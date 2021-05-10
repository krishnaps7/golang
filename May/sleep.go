package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	for {
		// z, err := os.Getwd()
		// fmt.Printf("Current path is: %+v\n", z)
		time.Sleep(2 * time.Second)
		data, err := ioutil.ReadFile("./test.txt")
		if err != nil {
			fmt.Printf("%+v\n", "SOmething wrong")
			return
		}
		fmt.Printf("%+v\n", string(data))
	}
}
