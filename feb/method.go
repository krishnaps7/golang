package main

import "fmt"

type location float64

func (c location) square() location {
	// fmt.Println(c * c)
	return (c * c)
}
func main() {
	var bita location = 20.023
	data := bita.square()
	fmt.Println(data)
}
