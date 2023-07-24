package main

import (
	"fmt"
	"flag"
)

func main() {
	flag.Parse()
	args := flag.Args()

	fmt.Println("Output", args)
	fmt.Println("Args", flag.Args())
}
