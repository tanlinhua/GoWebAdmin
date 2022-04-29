package utils

import (
	"github.com/lithammer/shortuuid/v3"
	gonanoid "github.com/matoous/go-nanoid/v2"
	go_uuid "github.com/satori/go.uuid"
)

func UUID() string {
	u2 := go_uuid.NewV4()
	return u2.String()
}

func ShortUUID() string {
	return shortuuid.New()
}

// JS版本：https://github.com/ai/nanoid
func Nanoid() string {
	id, err := gonanoid.New() // Simple usage
	// id, err = gonanoid.New(5) // Custom length
	// id, err = gonanoid.Generate("abcdefg", 10) // Custom alphabet
	if err == nil {
		return id
	}
	return ""
}
