package utils

import "github.com/google/uuid"

//GetUUID returns a new UUID
func GetUUID() uuid.UUID {
	return uuid.New()
}
