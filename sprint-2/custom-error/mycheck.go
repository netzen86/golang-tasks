//go:build !solution

package mycheck

import (
	"errors"
	"fmt"
	"regexp"
)

type CheckError []error

func (se CheckError) Error() error {
	var result error
	for i, v := range se {
		if i == 0 {
			result = v
			continue
		}
		result = fmt.Errorf("%w;%w", result, v)
	}
	return result
}

func MyCheck(input string) error {
	var result CheckError
	if regexp.MustCompile(`\d`).MatchString(input) {
		result = append(result, errors.New("found numbers"))
	}
	if len(input) >= 20 {
		result = append(result, errors.New("line is too long"))
	}
	if len(regexp.MustCompile(`\s`).FindAllString(input, -1)) != 2 {
		result = append(result, errors.New("no two spaces"))
	}
	return result.Error()
}
