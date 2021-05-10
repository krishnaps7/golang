// This will read file all at once
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	x, err := os.Getwd()
	fmt.Printf("%+v\n", x)
	os.Chdir("../")
	x, err = os.Getwd()
	fmt.Printf("%+v\n", x)
	if len(os.Args) != 2 {
		fmt.Println("please specify the path")
		return
	}
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error on the way: %+v\n", err)
		return
	}
	fmt.Printf("%+v\n", string(b))

}
