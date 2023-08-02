package main

import (
	"fmt"
	"go23/ex03/asm2/ulti"
)



func main() {
	word := "a123bc34d8ef34"
	number, arr := ulti.NumDifferentInteger(word)
	fmt.Println(number, arr)

}
