package main

import (
	"fmt"
	"go23/ex03/asm1/ulti"
)

func main() {
	arr := [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1},
		{0, 0, 0, 1, 1, 0},
	}

	count := ulti.CountRectangle(arr)
	fmt.Println(count) //6
}
