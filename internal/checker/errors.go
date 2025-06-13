package checker

import "fmt"

type UnreachableUrlError struct {
	Url string
	Err error
}

func (e *UnreachableUrlError) Error() string {
	return fmt.Sprintf("url inaccessible: %s: (%v)", e.Url, e.Err)
}

func (e *UnreachableUrlError) Unwrap() error {
	return e.Err
}
