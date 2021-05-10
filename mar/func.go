package main

import "fmt"

func agewhat(age int) string {
	if age > 20 {
		return "adult"
	} else {
		return "non-adult"
	}
}
func main() {
	var age = 21
	fmt.Println("You are an ", agewhat(age))
	var a string = '1'
	var b string = '2'
	c := a + b
	fmt.Printf("The type of c is %T and value is %v", c, c)
}
