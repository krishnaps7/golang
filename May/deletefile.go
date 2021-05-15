package main

import (
	"fmt"
	"os"
)

var filename = "testfile.txt"

func main() {
	if err := os.Remove(filename); err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("File has been deleted")

}
