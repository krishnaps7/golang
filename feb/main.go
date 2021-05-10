package main

import (
	"fmt"
	// "math/rand"
	"strings"
	"time"
)

func main() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Printf("%+v\n", "liftoff")
	fmt.Println("hello, welcome to go")
	var name = "krishna sabaris Pattamsetty"
	var status = strings.Contains(name, "krishna")
	fmt.Println("Name found:", status)
	// var number = rand.Intn(10)+1
	// fmt.Printf("Random number is %v\n", number)
	const age = 40
	// var minor = age < 18
	fmt.Println("Is minor? ", age > 18)
	// var data ="this is imp data"
	// fmt.Printf("data contains key word?%+v\n",strings.Contains(data,"imp"))
	// fmt.Printf("My age is %+v, am I a minor?%v\n",age, minor)
	var Dec = true
	var Winter = true
	var Africa = false

	if !Dec || !Winter || !Africa {
		fmt.Println("It's not cold time")
	}
}
