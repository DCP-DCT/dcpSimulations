package main

import (
	"github.com/DCP-DCT/DCP"
	"math/rand"
)

// EstablishNodeRelationships takes an array of CtNodes and
// randomly sets the CtNode.ReachableNodes attribute on the
// supplied nodes.
func EstablishNodeRelationships(nodes []*DCP.CtNode) {
	if len(nodes) < 2 {
		return
	}

	for i := 0; i < len(nodes); i++ {
		current := nodes[i]

		numbersToAdd := rand.Intn(len(nodes))

		if numbersToAdd == len(nodes) {
			numbersToAdd = numbersToAdd - 1
		}
		for j := 0; j < numbersToAdd; j++ {
			// randomNodeIndex between {0, len(nodes)}
			var randomNodeIndex int
			it := 1
			for {
				it++
				randomNodeIndex = rand.Intn(len(nodes))
				if (!contains(current.ReachableNodes, nodes[randomNodeIndex].Channel)) && (current.Id != nodes[randomNodeIndex].Id) {
					break
				}
				if it > 10 {
					break
				}
			}

			// Failed to find random index
			if it > 10 {
				continue
			}

			reachableNode := nodes[randomNodeIndex]

			// Don't add self
			if reachableNode.Id == current.Id {
				continue
			}

			current.ReachableNodes = append(current.ReachableNodes, reachableNode.Channel)
		}
	}
}

func contains(source []chan *DCP.CalculationObjectPaillier, target chan *DCP.CalculationObjectPaillier) bool {
	for _, c := range source {
		if c == target {
			return true
		}
	}

	return false
}
