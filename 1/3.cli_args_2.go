package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	s, sep := "", ""
	for _,arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}

// go run 3.cli_args_2.go hellooooo how are you