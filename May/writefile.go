package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("%+v\n", "PLease enter file and content for it")
		return
	}
	if err := ioutil.WriteFile(os.Args[1], []byte(os.Args[2]), 0644); err != nil {
		fmt.Printf("Error: %+v\n", err)
	}
}
