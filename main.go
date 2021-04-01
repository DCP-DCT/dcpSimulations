package main

import (
	"fmt"
	"github.com/DCP-DCT/DCP"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
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

	//os.Stdout = os.Stderr
	os.Stdout = nil
	runSimulation(100, false, false)
}

func redirectStderr(f *os.File) {
	err := syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		log.Fatalf("Failed to redirect stderr to file: %v", err)
	}
}

func createNodes(numberOfNodes int, config DCP.CtNodeConfig) []*DCP.CtNode {
	var nodes []*DCP.CtNode
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numberOfNodes; i++ {
		node := DCP.NewCtNode(GenerateIdTable(rand.Intn(25-1)+1), config)

		e := node.Co.KeyGen()
		if e != nil {
			fmt.Println(e.Error())
			break
		}

		nodes = append(nodes, node)
	}

	return nodes
}

func runSimulation(numberOfNodes int, randomStart bool, cluster bool) {
	td := 1 * time.Millisecond
	clusterSize := 3

	config := DCP.NewCtNodeConfig()
	config.NodeVisitDecryptThreshold = 3
	config.SuppressLogging = false
	config.Throttle = &td
	config.CoTTL = 10

	nodes := createNodes(numberOfNodes, config)

	if cluster {
		EstablishNodeRelationshipsLocalClusters(nodes, clusterSize)
	} else {
		EstablishNodeRelationShipAllInRange(nodes)
	}

	lock1 := &sync.RWMutex{}

	closeMonitor := make(chan struct{})
	go LaunchMonitor(nodes, closeMonitor, lock1, true)

	stop := make(chan struct{})
	for _, node := range nodes {
		node.Listen()

		if randomStart {
			go func(node *DCP.CtNode) {
				RandomCalculationProcessInitiator(node, stop)
			}(node)
		} else {
			go func(node *DCP.CtNode) {
				CalculationProcessInitiator(node)
			}(node)
		}
	}

	for {
		select {
		case <-closeMonitor:
			e := generateReport(nodes, cluster)
			if e != nil {
				panic(e)
			}

			close(stop)

			return
		default:
			nGoR := runtime.NumGoroutine()
			nBranches := DCP.GetNrOfActiveBranches()
			fmt.Println(nGoR, nBranches)
		}
	}
}
