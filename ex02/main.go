package main

import (
	"flag"
	"fmt"
	"go23/ex02/ulti"
)

func main() {
	intDataTypePtr := flag.Bool("int", false, "Specify if the data type is int")
	floatDataTypePtr := flag.Bool("float", false, "Specify if the data type is float")
	stringDataTypePtr := flag.Bool("string", false, "Specify if the data type is string")
	mixDataTypePtr := flag.Bool("mix", false, "Specify if the data type is mix")
	flag.Parse()

	// Determine the selected data type
	dataType := ""
	if *intDataTypePtr {
		dataType = "int"
	} else if *floatDataTypePtr {
		dataType = "float"
	} else if *stringDataTypePtr {
		dataType = "string"
	} else if *mixDataTypePtr {
		dataType = "mix"
	} else {
		fmt.Println("No data type selected. Please use one of the following flags: -int, -float, -string, -mix.")
		return
	}

	parts := flag.Args()
	values, err := ulti.ParseInput(dataType, parts)
	if err != nil {
		fmt.Println(err)
	} else {
		result, err := ulti.Sort(dataType, values)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}
