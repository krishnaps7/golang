package main

import (
	"fmt"
	"os"
)

func main() {
	var filename = "newfile.txt"
	if err := os.Rename(filename, "testfile-renamed.txt"); err != nil {
		fmt.Println("Error:", err)
	}
}
