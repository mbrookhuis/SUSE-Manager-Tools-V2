// Package uuid - generate unique UUID
package uuid

import guuid "github.com/google/uuid"

// GenerateUniqueID  generate unique UUID
//
// return: string unique id
func GenerateUniqueID() string {
	return guuid.New().String()
}
