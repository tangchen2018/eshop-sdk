package utils

import (
	"errors"
	"time"
)

func Err(e string) error {
	return errors.New(e)
}

func PString(e string) *string {
	return &e
}

func TimestampSecond() int64 {
	return time.Now().UnixMilli() / 1000
}
