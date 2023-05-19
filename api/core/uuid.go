package core

import "github.com/google/uuid"

// GenerateID creates a new UUID
func GenerateID() string {
	return uuid.New().String()
}
