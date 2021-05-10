package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// z, err := os.Getwd()
	// if err != nil {
	// 	fmt.Printf("%+v\n", err)
	// 	return
	// }
	// fmt.Printf("%+v\n", z)
	// if err := os.Chdir("/"); err != nil {
	// 	fmt.Printf("%+v\n", err)
	// }
	// z, err = os.Getwd()
	// if err != nil {
	// 	fmt.Printf("%+v\n", err)
	// 	return
	// }
	// fmt.Printf("%+v\n", z)
	// fmt.Printf("%+v\n", filepath.Dir("~/Documents/GCP"))
	// os.Chdir("feb")
	d, err := os.Getwd()
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	fmt.Printf("%+v\n", d)
	data, err := filepath.Glob("*go")
	fmt.Printf("%+v\n", data)

	if len(os.Args) != 2 {
		fmt.Printf("%+v\n", "please specify a path")
		return
	}
	root, err := filepath.Abs(os.Args[1])
	if err != nil {
		fmt.Printf("%+v\n", "Can't get absolute path")
		return
	}
	fmt.Println("Returning files in the path", root)
	var c struct {
		file int
		dir  int
	}
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			c.dir++
		} else {
			c.file++
		}
		fmt.Println("-", info.ModTime())
		return nil
	})
	fmt.Printf("Total: %d files and %d directories\n", c.file, c.dir)
}
