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
			ReachableNodes: nil,
			Channel:        make(chan *DCP.CalculationObjectPaillier),
		}
		e := node.Co.KeyGen()
		if e != nil {
			fmt.Println(e.Error())
			break
		}

		node.Listen()
		nodes = append(nodes, node)
	}

	EstablishNodeRelationships(nodes)

	initialNode := nodes[0]

	e := DCP.InitRoutine(DCP.PrepareIdLenCalculation, initialNode)
	if e != nil {
		fmt.Println(e)
	}

	initialNode.Broadcast()

	fmt.Scanln()
}
