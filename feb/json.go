package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		Lat  float64 `json:"Latitude"`
		Long float64 `json:"Longitude"`
	}
	curiosity := location{-4.2301, 137.4418}
	bytes, err := json.Marshal(curiosity)
	ExitOnError(err)
	fmt.Println(string(bytes))
}
func ExitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
