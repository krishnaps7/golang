package main

import (
	"fmt"
	"time"
)

// "math/rand"
func test() {

	future := time.Unix(12364903543, 0)
	fmt.Printf("%+v\n", future)
	var pi rune = 960
	var alpha rune = 940
	var bang byte = 33
	var grade = 'A'
	var data = "some data given"
	fmt.Printf("%+v %+v %+v \n", pi, alpha, bang)
	fmt.Printf("%+c %+c %+c %c\n", pi, alpha, bang, grade+4)
	fmt.Println(len(data), "bytes")
}
func main() {
	// var count = 0
	// for count = 10; count > 0; count-- {
	// 	fmt.Println(count)
	// }
	// for c2 := 20; c2 > 10; c2-- {
	// 	fmt.Println(c2)
	// }
	// if num := rand.Intn(3); num == 0 {
	// 	fmt.Println("num is zero")
	// } else {
	// 	fmt.Println("the num is", num)
	// }
	// var f64 float64 = 1.0 / 3
	// fmt.Println(f64)
	// fmt.Printf("%v\n", f64)
	test()
}
