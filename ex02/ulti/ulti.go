package ulti

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
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
func ParseInput(parts []string) (string, []interface{}, error) {
	if len(parts) < 2 {
		return "", nil, errors.New("wrong input")
	}
	dataType := parts[0]
	rawvalue := parts[1:]
	var values []interface{}
	for _, v := range rawvalue {
		val := parseValue(v)
		values = append(values, val)
	}
	return dataType, values, nil
}
func parseValue(value string) interface{} {
	num, err := strconv.Atoi(value)
	if err == nil {
		return num
	}
	return value
}
func Sort(dataType string, values []interface{}) ([]interface{}, error) {
	sort.Sort(sortableValues(values))
	return values, nil
}
