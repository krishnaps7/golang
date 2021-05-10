package main

import "fmt"

func main() {
	// m := map[string]int
	m := make(map[string]int) //
	m["age"] = 12
	fmt.Printf("%+v\n", m)
}
