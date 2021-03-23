package main

import (
	"fmt"
	"github.com/DCP-DCT/DCP"
	"math/rand"
	"os"
	"time"
)

func main() {
	//benchmarkEncryption(50)

	temp := os.Stdout
	os.Stdout = nil
	runSimulation(9)
	os.Stdout = temp
}

func createNodes(numberOfNodes int, config *DCP.CtNodeConfig) []*DCP.CtNode {
	var nodes []*DCP.CtNode
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numberOfNodes; i++ {
		node := DCP.NewCtNode(GenerateIdTable(rand.Intn(25)), config)

		e := node.Co.KeyGen()
		if e != nil {
			fmt.Println(e.Error())
			break
		}

		nodes = append(nodes, node)
	}

	return nodes
}

func benchmarkEncryption(numberOfNodes int) {
	config := &DCP.CtNodeConfig{
		NodeVisitDecryptThreshold: 5,
	}

	nodes := createNodes(numberOfNodes, config)

	for _, node := range nodes {
		node.Listen()
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

	fmt.Printf("Initial Node Counter %d, Node Cipher %s\n", initialNode.Co.Counter, msg.String())
}

func runSimulation(numberOfNodes int) {
	td := 10 * time.Millisecond

	config := DCP.NewCtNodeConfig()
	config.NodeVisitDecryptThreshold = 2
	config.SuppressLogging = true
	config.Throttle = &td

	nodes := createNodes(numberOfNodes, config)
	EstablishNodeRelationShipAllInRange(nodes)

	for _, node := range nodes {
		node.Listen()
	}

	closeMonitor := make(chan struct{})
	go LaunchMonitor(nodes, closeMonitor)

	stop := make(chan struct{})
	for _, node := range nodes {
		go func(node *DCP.CtNode) {
			RandomCalculationProcessInitiator(node, stop)
		}(node)
	}

	for {
		select {
		case <-closeMonitor:
			close(stop)
			return
		}
	}
}
