package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	fmt.Println("Enter text here")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "" {
			break	
		}
		counts[input.Text()]++
	}
	fmt.Println("text Counter ", counts)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}