package main

import "fmt"

type Human struct {
	name  string
	phone string
	age   int
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) sayHi() {
	fmt.Printf("From Human database: My name is %s, call me on %s", h.name, h.phone)
}

func (e *Employee) sayHi() {
	fmt.Printf("From Employee database: My name is %s, call me on %s", e.name, e.phone)
}

func main() {
	sam := Employee{Human{"sam", "888-999-777", 27}, "local company"}
	mark := Student{Human{"mark", "000-111-222", 20}, "local school"}
	sam.sayHi()
	fmt.Println()
	mark.sayHi()
}
