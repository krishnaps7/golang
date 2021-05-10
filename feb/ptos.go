package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}

func main() {
	krishna := person{
		name:       "krishna",
		superpower: "Imagination",
		age:        31,
	}
	birthday(&krishna)
	fmt.Printf("%+v\n", krishna)
}
