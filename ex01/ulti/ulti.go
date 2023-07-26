package ulti

import (
	"errors"
	"strings"
)

func Convert(input string) string {
	parts := strings.Split(input, " ")
	handleMiddleName(parts)
	handleMiddleName(parts)
	countryCode, firstName, middleName, lastName, err := handleMiddleName(parts)
	if err != nil {
		return err.Error()
	}
	fullName := handleCountryCode(countryCode, firstName, middleName, lastName)
	return fullName
}

func handleMiddleName(parts []string) (string, string, string, string, error) {
	firstName := parts[0]
	var lastName string
	var middleName string
	var countryCode string
	if len(parts) >= 3 {
		countryCode = strings.TrimSpace(parts[len(parts)-1])
		middleName = strings.Join(parts[1:len(parts)-2], " ")
		lastName = parts[len(parts)-2]
	} else if len(parts) == 2 {
		lastName = parts[1]
		countryCode = strings.TrimSpace(parts[0])
	} else {
		return "", "", "", "", errors.New("not enough parts to form a name")
	}
	return countryCode, firstName, middleName, lastName, nil

}
func handleCountryCode(countryCode, firstName, middleName, lastName string) string {
	var fullName string
	switch countryCode {
	case "VN", "CN", "JP", "KR", "IN", "TH", "ID", "PH", "MY", "SG":
		if middleName != "" {
			fullName = lastName + " " + middleName + " " + firstName
		} else {
			fullName = lastName + " " + firstName
		}
	case "US", "CA", "GB", "FR", "DE", "IT", "ES", "AU", "NZ", "BR":
		if middleName != "" {
			fullName = firstName + " " + middleName + " " + lastName
		} else {
			fullName = firstName + " " + lastName
		}
	default:
		return "Invalid contry code"
	}
	return fullName
}
