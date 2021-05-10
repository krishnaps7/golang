package main

import "fmt"

func main() {
	type mystring string
	const Greeting = "hello"
	var s1 string = Greeting
	var s2 mystring = Greeting
	fmt.Printf("%+v %T\n%v %T\n", s1, s1, s2, s2)
}
