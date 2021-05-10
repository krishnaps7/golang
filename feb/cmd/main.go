package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	ex, _ := exec.LookPath("kubectl")
	cmd := &exec.Cmd{
		Path: ex,
		Args: []string{ex, "get", "pods", "--context", "dev-eu-west"},
		// Args:   []string{ex, "version"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
	fmt.Println(cmd) // It just prints the cmd but doesn't execute it
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
