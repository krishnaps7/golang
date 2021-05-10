package main

import "fmt"

func main() {
	planets := []string{"Earth", "Mars", "Venus"}
	newPlanets := terraform("new", planets...)
	fmt.Println(newPlanets)
}

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))
	for i, j := range worlds {
		fmt.Println(i, j)
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}
