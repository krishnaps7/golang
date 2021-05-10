package main

import (
	"fmt"
	"math"
)

func main() {
	temparature := map[string]int{
		"Earth": 15,
		"Venus": 464,
	}
	// temp := temparature["Earth"]
	temparature["Venus"] = 200
	temparature["Mars"] = 220
	fmt.Println(temparature)
	fmt.Println("The value for moon is:", temparature["Moon"])
	if moon, ok := temparature["Venus"]; ok {
		fmt.Println("Venus's temparature is: ", moon)
	} else {
		fmt.Println("Where is the moon data?")
	}
	tmp := []float64{-25.0, 34.2, 34.2, 43.0, -234.1}
	frequency := make(map[float64]int)
	for _, t := range tmp {
		frequency[t]++
	}
	for t, num := range frequency {
		fmt.Printf("%+.2f occures %d times\n", t, num)
	}
	fmt.Println(frequency)
	group := make(map[float64][]float64)
	for _, i := range tmp {
		g := math.Trunc(i/10) * 10
		group[g] = append(group[g], i)
	}
	fmt.Println(group)
}
