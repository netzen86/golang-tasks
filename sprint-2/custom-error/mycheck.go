//go:build !solution

package mycheck

type CheckError struct {
	Errors []error
}

func (se CheckError) Error() []error {
	return se.Errors
}

func Error(text string) CheckError {
	var result CheckError

}

func MyCheck(input string) error {
	// TODO
	return nil
}
