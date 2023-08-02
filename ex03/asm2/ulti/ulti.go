package ulti

import "strconv"

func getNextPointer(word string, index, pointer int) int {
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
	return pointer
}
func NumDifferentInteger(word string) (int, []int) {
	var numarr []int
	index := 0
	for index < len(word) {
		pointer := index
		var target int
		pointer = getNextPointer(word, index, pointer)
		target, err := strconv.Atoi(string(word[index:pointer]))
		if err != nil {
			index++
			continue
		} else if numberExists(numarr, target) {
			index = pointer
			continue
		} else {
			index = pointer
			numarr = append(numarr, target)
		}
	}
	return len(numarr), numarr
}
func numberExists(arr []int, target int) bool {
	for _, num := range arr {
		if num == target {
			return true
		}
	}
	return false
}
