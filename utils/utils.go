package utils

import "github.com/google/uuid"
import "fmt"

//GetUUID returns a new UUID
func GetUUID() uuid.UUID {
	return uuid.New()
}

//GetDoubleUUID returns a new double length UUID
func GetDoubleUUID() string {
	return fmt.Sprintf("%s-%s\n", uuid.New().String(), uuid.New().String())
}
