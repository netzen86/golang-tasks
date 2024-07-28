//go:build !solution

package testequal

import (
	"fmt"
	"reflect"
)

const (
	i    string = "int"
	i8   string = "int8"
	i16  string = "int16"
	i32  string = "int32"
	i64  string = "int64"
	ui8  string = "uint8"
	ui16 string = "uint16"
	ui32 string = "uint32"
	ui64 string = "uint64"
	s    string = "string"
	ma   string = "map"
	sl   string = "slice"
	df   string = "NA"
)

func GetType(testValue interface{}) string {
	switch testValue.(type) {
	case int:
		return i
	case int8:
		return i8
	case int16:
		return i16
	case int32:
		return i32
	case int64:
		return i64
	case uint8:
		return ui8
	case uint16:
		return ui16
	case uint32:
		return ui32
	case uint64:
		return ui64
	case string:
		return s
	case map[string]string:
		return ma
	case []int:
		return sl
	case []byte:
		return sl
	default:
		return df
	}
}

func compare(expected, actual interface{}) bool {
	et, ea := GetType(expected), GetType(actual)
	switch {
	case et == ma && ea == ma || et == sl && ea == sl:
		return reflect.DeepEqual(expected, actual)
	case et != ea:
		return false
	case et == "NA" || ea == "NA":
		return false
	default:
		return expected == actual
	}
}

func createErrorText(data []interface{}) string {
	if len(data) == 3 {
		return fmt.Sprintf(data[0].(string), data[1].(int), data[2].(int))
	}
	if len(data) == 1 {
		return data[0].(string)
	}
	return ""
}

// AssertEqual checks that expected and actual are equal.
//
// Marks caller function as having failed but continues execution.
//
// Returns true iff arguments are equal.
func AssertEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	result := compare(expected, actual)
	if !result {
		t.Helper()
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
		t.Helper()
		t.Errorf(createErrorText(msgAndArgs))
	}
	return !result
	// panic("implement me")
}

// RequireEqual does the same as AssertEqual but fails caller test immediately.
func RequireEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	result := compare(expected, actual)
	if !result {
		t.Helper()
		t.Errorf(createErrorText(msgAndArgs))
		t.FailNow()
	}
	// panic("implement me")
}

// RequireNotEqual does the same as AssertNotEqual but fails caller test immediately.
func RequireNotEqual(t T, expected, actual interface{}, msgAndArgs ...interface{}) {
	result := compare(expected, actual)
	if result {
		t.Helper()
		t.Errorf(createErrorText(msgAndArgs))
		t.FailNow()
	}
}
