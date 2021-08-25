package utils

import (
	"github.com/lithammer/shortuuid/v3"
	go_uuid "github.com/satori/go.uuid"
)

func UUID() string {
	u2 := go_uuid.NewV4()
	return u2.String()
}

func ShortUUID() string {
	return shortuuid.New()
}
