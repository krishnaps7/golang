package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b = bytes.NewBuffer(make([]byte, 20))
	var text = []string{
		`data1data1data1data1data1data1data1`,
		`data2`,
		`data3`,
	}
	for i := range text {
		b.Reset()
		fmt.Printf("%+v\n", i)
		b.WriteString(text[i])
		fmt.Printf("len: %+v capacity: %v\n", b.Len(), b.Cap())
	}
	fmt.Printf("data: %+v\n", b)
}
