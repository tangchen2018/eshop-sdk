package utils

import "errors"

func Err(e string) error {
	return errors.New(e)
}

func PString(e string) *string {
	return &e
}
