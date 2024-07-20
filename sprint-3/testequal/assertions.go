//go:build !solution

package testequal

import (
	"fmt"
	"reflect"
)

func GetType(testValue interface{}) string {
	var resultType string
	switch testValue.(type) {
	case int:
		resultType = "int"
	case int8:
		resultType = "int8"
	case int16:
		resultType = "int16"
	case int32:
		resultType = "int32"
	case int64:
		resultType = "int64"
	case uint8:
		resultType = "uint8"
	case uint16:
		resultType = "uint16"
	case uint32:
		resultType = "uint32"
	case uint64:
		resultType = "uint64"
	case string:
		resultType = "string"
	case map[string]string:
		resultType = "map"
	case []int:
		resultType = "slice"
	case []byte:
		resultType = "slice"
	// case interface{}:
	// 	resultType = "interface {}"
	default:
		resultType = "NA"
	}
	return resultType
}

func compare(expected, actual interface{}) bool {
	et, ea := GetType(expected), GetType(actual)
	result := false
	switch {
	case et == "map" && ea == "map" || et == "slice" && ea == "slice":
		result = reflect.DeepEqual(expected, actual)
	case et != ea:
		result = false
	case et == "NA" || ea == "NA":
		result = false
	default:
		result = expected == actual
	}
	return result
}

func createErrorText(data []interface{}) string {
	result := ""
	if len(data) == 3 {
		result = fmt.Sprintf(data[0].(string), data[1].(int), data[2].(int))
	}
	if len(data) == 1 {
		result = data[0].(string)
	}
	return result
}

// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are equal.
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	result := compare(expected, actual)
	if !result {
		fmt.Println(createErrorText(msgAndArgs))
		t.Errorf(createErrorText(msgAndArgs))
	}
	return result
	// panic("implement me")
}

// AssertNotEqual checks that expected and actual are not equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are not equal.
func AssertNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	result := compare(expected, actual)
	if result {
		t.Errorf(createErrorText(msgAndArgs))
	}
	return !result
	// panic("implement me")
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	result := compare(expected, actual)
	if !result {
		t.Errorf(createErrorText(msgAndArgs))
		t.FailNow()
	}
	// panic("implement me")
}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	result := compare(expected, actual)
	if result {
		t.Errorf(createErrorText(msgAndArgs))
		t.FailNow()
	}
}
