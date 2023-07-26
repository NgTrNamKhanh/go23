package ulti

import (
	//"errors"
	"fmt"
	"strings"
	//"reflect"
	"sort"
	"strconv"
)

type sortableValues []interface{}

func (sv sortableValues) Len() int      { return len(sv) }
func (sv sortableValues) Swap(i, j int) { sv[i], sv[j] = sv[j], sv[i] }
func (sv sortableValues) Less(i, j int) bool {
	switch sv[i].(type) {
	case int:
		return sv[i].(int) < sv[j].(int)
	case float64:
		return sv[i].(float64) < sv[j].(float64)
	case string:
		return sv[i].(string) < sv[j].(string)
	default:
		// Add custom handling for other data types if needed
		return false
	}
}

func ParseInput(dataType string, parts []string) ([]interface{}, error) {
	var values []interface{}
	for _, v := range parts {
		val, err := parseValue(dataType, v)
		if err != nil {
			return nil, err
		}
		values = append(values, val)
	}
	return values, nil
}
func parseValue(dataType, value string) (interface{}, error) {
	switch dataType {
	case "int":
		num, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("invalid integer value: %s", value)
		}
		return num, nil

	case "float":
		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid float value: %s", value)
		}
		return num, nil

	default:
		return value, nil
	}
}

func Sort(dataType string, values []interface{}) (string, error) {
	sort.Sort(sortableValues(values))
	return joinValues(values), nil
}
func joinValues(values []interface{}) string {
	var result []string
	for _, v := range values {
		result = append(result, fmt.Sprintf("%v", v))
	}
	return strings.Join(result, " ")
}

