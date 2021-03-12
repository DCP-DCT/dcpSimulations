package main

import "github.com/google/uuid"

func GenerateIdTable(length int) []string {
	ids := make([]uuid.UUID, length)

	var idsAsString []string
	for _, id := range ids {
		id = uuid.New()
		idsAsString = append(idsAsString, id.String())
	}

	return idsAsString
}
