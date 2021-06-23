package main

import (
	"flag"
	"fmt"
)

func main(){
	// Raw flags example
	a := flag.String("arg", "none", "example flag")
	flag.Parse()

	fmt.Println(*a)
}
