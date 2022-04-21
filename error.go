package maps

import "fmt"

type Error[Key comparable, Val any] struct {
	error
	Key Key
	Val Val
}

func (e *Error[Key, Val]) Unwrap() error {
	return e.error
}

func (e *Error[Key, Val]) Error() string {
	return fmt.Sprintf("error processing key `%v`: %s", e.Key, e.error.Error())
}
