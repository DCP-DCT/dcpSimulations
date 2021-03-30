package main

import (
	"fmt"
	"github.com/DCP-DCT/DCP"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

func main() {
	//benchmarkEncryption(50)
	fName := "dcp-sim-error-log-" + strconv.Itoa(int(time.Now().UnixNano())) + ".txt"
	fp := filepath.Join("log", fName)
	f, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		panic(err.Error())
	}

	defer f.Close()

	redirectStderr(f)

	runSimulation(40)
}

func redirectStderr(f *os.File) {
	err := syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		log.Fatalf("Failed to redirect stderr to file: %v", err)
	}
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
	config.NodeVisitDecryptThreshold = 5
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
			e := generateReport(nodes)
			if e != nil {
				panic(e)
			}

			close(stop)

			return
		}
	}
}
