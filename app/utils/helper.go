package utils

import (
	uuid "github.com/satori/go.uuid"
)

// GenerateUUID
func GenerateUUID() string {
	myuuid := uuid.NewV4()
	return myuuid.String()
}
