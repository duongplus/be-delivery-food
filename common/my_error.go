package common

import (
	"errors"
	"log"
)

var (
	RecordNotFound       = errors.New("record not found")
	ErrNameCannotBeEmpty = NewCustomError(nil, "restaurant name can't be blank", "ErrNameCannotBeEmpty")
	ErrFileTooLarge      = NewCustomError(
		errors.New("file too large"),
		"file too large",
		"ErrFileTooLarge",
	)
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
