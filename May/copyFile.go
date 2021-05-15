package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var src = "testfile-renamed.txt"
	var dst = "copied.txt"
	copyFile(src, dst)
}
func copyFile(from, to string) (int64, error) {
	src, err := os.Open(from)
	if err != nil {
		fmt.Println("Error: ", err)
		return 0, err
	}
	defer src.Close()
	dst, err := os.OpenFile(to, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
		return 0, err
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
