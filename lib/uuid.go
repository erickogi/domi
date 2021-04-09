package lib

import "github.com/google/uuid"

func getUUID() string {
	return uuid.NewString()
}
