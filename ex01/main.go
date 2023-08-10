package main

import (
	"fmt"
	"github.com/NgTrNamKhanh/ex01/ulti"
	"os"
	"strings"
)

func main() {
	input := strings.Join(os.Args[1:], " ")
	result := ulti.Convert(input)
	fmt.Print("Output:", result)
}
