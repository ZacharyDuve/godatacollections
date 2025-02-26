package godatacollections

import "errors"

const EMPTY_ERROR_MSG string = "empty error"

var emptyError error

func init() {
	emptyError = errors.New(EMPTY_ERROR_MSG)
}

func EmptyError() error {
	return emptyError
}

func IsEmptyError(err error) bool {
	// Since TODO
	return err == emptyError
}
