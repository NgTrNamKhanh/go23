package main

import (
	"fmt"
	"strconv"
)

func numDifferentInteger(word string) []int {
	var numarr []int
	index := 0
	for index < len(word)-1 {
		pointer := index
		var target int
		for i := index; i <= len(word); i++ {
			if i == len(word) {
				pointer = i
				break
			}
			_, err := strconv.Atoi(string(word[i]))
			pointer = i
			if err != nil {
				break
			} else {
				continue
			}
		}
		target, err := strconv.Atoi(string(word[index:pointer]))
		if numberExists(numarr, target) || err != nil {
			index++
			continue
		} else {
			index = pointer
			numarr = append(numarr, target)
		}
	}
	return numarr
}
func numberExists(arr []int, target int) bool {
	for _, num := range arr {
		if num == target {
			return true
		}
	}
	return false
}

func main() {
	word := "A1b01c001"
	arr := numDifferentInteger(word)
	fmt.Println(len(arr), arr)

}
