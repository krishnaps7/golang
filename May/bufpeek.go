package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// z, err := os.Getwd()
	// fmt.Printf("%+v\n", z)
	if len(os.Args) != 2 {
		fmt.Printf("%+v\n", "Please enter the path to file")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	var rowCount int
	for err == nil {
		var b []byte
		for moar := true; err == nil && moar; {
			b, moar, err = r.ReadLine()
			if err == nil {
				fmt.Printf("data: %+v\n", string(b))
			}
		}
		if err == nil {
			fmt.Println()
			rowCount++
		}
	}
	if err != nil && err != io.EOF {
		fmt.Printf("Error: %+v\n", err)
		return
	}
	fmt.Printf("Row count: %+v\n", rowCount)
}
