package ulti

import (
	//"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type sortableValues []interface{}

func (sv sortableValues) Len() int      { return len(sv) }
func (sv sortableValues) Swap(i, j int) { sv[i], sv[j] = sv[j], sv[i] }
func (sv sortableValues) Less(i, j int) bool {
	return fmt.Sprintf("%v", sv[i]) < fmt.Sprintf("%v", sv[j])
}

func TypeCheck(dataType string, values []string) bool {
	return dataType == reflect.TypeOf(values[0]).String()
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
	result := interfacetostring(values)
	return result, nil
}
func interfacetostring(values []interface{}) string {
	stringarr := make([]string, len(values))
	for _, val := range values {
		stringval := strconv.Itoa(val.(int))
		stringarr = append(stringarr, stringval)
	}
	result := strings.Join(stringarr, " ")
	return result
}
