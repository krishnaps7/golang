package main

import "fmt"

func main() {
	dwarf := [8]string{"earth", "sun", "saturn", "Jupitar"}
	for i := 0; i < len(dwarf); i++ {
		fmt.Println(dwarf[i])
	}
	for i := range dwarf {
		fmt.Println("New " + dwarf[i])
	}
	var board [8][8]string
	board[0][0] = "p"
	board[0][7] = "p"
	for col := range board[1] {
		board[1][col] = "p"
	}
	fmt.Print(board)
	fmt.Println(board[1])
}
