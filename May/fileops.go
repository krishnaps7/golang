package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var filename = "testfile.txt"

func main() {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer f.Close()
	data := "some data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file nowsome data is generated and will be written to the file now"
	if err = ioutil.WriteFile(filename, []byte(data), 0644); err != nil {
		fmt.Println("Error", err)
		return
	}
	if err = os.Truncate(filename, 500); err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
