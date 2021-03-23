package main

import (
	"github.com/DCP-DCT/DCP"
	"github.com/google/uuid"
)

type ContributionRecord struct {
	NodeId  uuid.UUID
	Updates map[DCP.ControlEntity]int
}

func GenerateIdTable(length int) []string {
	ids := make([]uuid.UUID, length)

	var idsAsString []string
	for _, id := range ids {
		id = uuid.New()
		idsAsString = append(idsAsString, id.String())
	}

	return idsAsString
}
