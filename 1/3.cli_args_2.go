package main

import (
	"fmt"
	"os"
)

func main(){
	s, sep := "", ""
	for _,arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// go run 3.cli_args_2.go hellooooo how are you