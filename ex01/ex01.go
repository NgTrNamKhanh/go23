package main

import (
	"bufio"
	"fmt"
	"go23/ex01/ulti"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Standard input

	fmt.Print("Enter your input: ")

	input, err := reader.ReadString('\n')

	result := ulti.Convert(input,err);
	fmt.Print(result)
}
