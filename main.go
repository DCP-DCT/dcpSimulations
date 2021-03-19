package main

import (
	"fmt"
	"github.com/DCP-DCT/DCP"
	"github.com/google/uuid"
	"math/rand"
	"os"
	"time"
)

func main() {
	benchmarkEncryption(50)
}

func benchmarkEncryption(numberOfNodes int) {
	temp := os.Stdout
	os.Stdout = nil


	var nodes []*DCP.CtNode
	rand.Seed(time.Now().UnixNano())

	config := &DCP.CtNodeConfig{
		NodeVisitDecryptThreshold: 5,
	}

	for i := 0; i < numberOfNodes; i++ {
		node := &DCP.CtNode{
			Id:           uuid.New(),
			Co:           &DCP.CalculationObjectPaillier{
				TransactionId:        uuid.New(),
				Counter:   0,
			},
			Ids:          GenerateIdTable(rand.Intn(25)),
			HandledCoIds: make(map[uuid.UUID]struct{}),
			TransportLayer: &DCP.ChannelTransport{
				DataCh:         make(chan *[]byte),
				StopCh:         make(chan struct{}),
				ReachableNodes: make(map[chan *[]byte]chan struct{}),
			},
			Config: config,
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
	EstablishNodeRelationShipAllInRange(nodes)

	e := DCP.InitRoutine(DCP.PrepareIdLenCalculation, initialNode)
	if e != nil {
		fmt.Println(e)
	}

	initialNode.Broadcast(nil)

	time.Sleep(10 * time.Second)
	msg := initialNode.Co.Decrypt(initialNode.Co.Cipher)

	os.Stdout = temp
	fmt.Printf("Initial Node Counter %d, Node Cipher %s\n", initialNode.Co.Counter, msg.String())
}
