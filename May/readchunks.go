package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("%+v\n", "Enter path to file")
		return
	}
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		return
	}
	defer f.Close()
	var b = make([]byte, 16)
	for n := 0; err == nil; {
		// fmt.Println("n: ", n, "err: ", err)
		n, err = f.Read(b)
		// fmt.Println("2st n: ", n)
		if err == nil {
			// fmt.Printf("%+v\n", "Inside if block")
			// fmt.Print(string(b[:n-2]), "\n")
			// fmt.Print(string(b[:n-1]), "\n")
			fmt.Print(string(b[:n]), "\n")
			// fmt.Print(string(b), "\n")
			fmt.Println("------------------------")
		}
		// time.Sleep(time.Second * 2)
		// fmt.Println("n: ", n, "err: ", err)
	}
	if err != nil && err != io.EOF {
		fmt.Printf("Error: %+v\n", err)
	}
}
