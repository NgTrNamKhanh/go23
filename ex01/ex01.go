package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/NgTrNamKhanh/go23/ex01/ulti/ulti"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Standard input

	fmt.Print("Enter your input: ")

	input, err := reader.ReadString('\n')


	result := 
	fmt.Print()

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	parts := strings.Split(input, " ")

	firstName := parts[0]
	var lastName string
	var middleName string

	var countryCode string
	if len(parts) >= 3 {
		countryCode = strings.TrimSpace(parts[len(parts)-1])  // Trim the newline character
		middleName = strings.Join(parts[1:len(parts)-2], " ") // Fix this line
		lastName = parts[len(parts)-2]
	} else if len(parts) == 2 {
		lastName = parts[1]
		countryCode = strings.TrimSpace(parts[0]) // Trim the newline character
	} else {
		fmt.Println("Error")
		return
	}

	var fullName string
	fmt.Println(countryCode)
	switch countryCode {
	case "VN", "CN", "JP", "KR", "IN", "TH", "ID", "PH", "MY", "SG":
		fullName = lastName + " " + middleName + " " + firstName
	case "US", "CA", "GB", "FR", "DE", "IT", "ES", "AU", "NZ", "BR":
		fullName = firstName + " " + middleName + " " + lastName
	default:
		fmt.Println("Invalid country code")
		return
	}

	fmt.Println(fullName)
}


