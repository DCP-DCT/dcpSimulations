package main

import (
	"fmt"
	"github.com/DCP-DCT/DCP"
	"github.com/google/uuid"
	"math/rand"
)

func main() {
	benchmarkEncryption(10)
}

func benchmarkEncryption(numberOfNodes int) {
	var nodes []*DCP.CtNode

	for i := 0; i < numberOfNodes; i++ {
		node := &DCP.CtNode{
			Id:             uuid.New(),
			Co:             &DCP.CalculationObjectPaillier{},
			Ids:            GenerateIdTable(rand.Intn(25)),
			ReachableNodes: make(map[chan *DCP.CalculationObjectPaillier]struct{}),
			Channel:        make(chan *DCP.CalculationObjectPaillier),
			HandledCoIds:   make(map[uuid.UUID]struct{}),
		}
		e := node.Co.KeyGen()
		if e != nil {
			fmt.Println(e.Error())
			break
		}

		node.Listen()
		nodes = append(nodes, node)
	}

	initialNode := nodes[0]
	EstablishNodeRelationships(nodes, initialNode)

	e := DCP.InitRoutine(DCP.PrepareIdLenCalculation, initialNode)
	if e != nil {
		fmt.Println(e)
	}

	initialNode.Broadcast(nil)

	fmt.Scanln()
}
