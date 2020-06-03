package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	pathList := filepath.SplitList(os.Getenv("PATH"))
	query := os.Args[1:]

	if !(len(os.Args) > 1) {
		fmt.Println("No Search term.")
		return
	}

	fmt.Println()
	for _, q := range query {
		for i, pi := range pathList {
			if strings.Contains(pi, q) {
				fmt.Printf("%-2d: %s\n", i+1, pi)
			}
		}
	}
	fmt.Println()
}
