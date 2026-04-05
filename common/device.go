package common

import "github.com/google/uuid"

func GenerateDeviceID() string {
	return uuid.New().String()
}
