package main

import "fmt"

// type akshara string
// type fun func(akshara) akshara
// type n [2]int

// func main() {
// 	var count n
// 	a := new(int)
// 	// var m [2]int = int(n)
// 	var i interface{} = "hello"
// 	z, ok := i.(int)
// 	fmt.Printf("%+v %v\n", z, ok)
// 	fmt.Printf("%v %T %v\n", a, a, *a)
// 	fmt.Printf("%+v\n", count)
// 	count = n{10, 20}
// 	fmt.Printf("%+v\n", count)
// 	var name akshara
// 	var description akshara
// 	name = "krishna"
// 	description = "some data given here"
// 	var f fun = func(x akshara) akshara {
// 		return akshara("Welcome ") + x
// 	}
// 	fmt.Printf("%+v\n", f(name))
// 	// someOther(f)
// 	fmt.Println(name)
// 	fmt.Println(description)
// 	var b interface{} = 10
// 	switch b.(type) {
// 	case int:
// 		fmt.Printf("%+v\n", "b is an int")
// 	case string:
// 		fmt.Printf("%+v\n", "b is a string")

// 	}
// }

// someOther(fn fun){
// 	data := fn()
// 	fmt.Sprintf("%s",data)
// }
func main() {
	var m = 10
	for n := true; m > 0 && n; {
		fmt.Printf("n: %+v, m:%v\n", n, m)
		m--
		if m == 4 {
			n = false
		}
	}
}
