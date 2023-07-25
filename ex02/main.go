package main

import (
	"flag"
	"fmt"
	"go23/ex02/ulti"
)

func main() {
	flag.Parse()
	parts := flag.Args()
	dataType, values, err := ulti.ParseInput(parts)
	if err != nil {
		fmt.Println(err)
	} else {
		//if ulti.TypeCheck(dataType, values) {
		result, err := ulti.Sort(dataType, values)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
		//} else {
		//fmt.Println("Wrong type")
		//}
	}

}
