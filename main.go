package main

import (
	"fmt"
	"github.com/DCP-DCT/DCP"
	"github.com/ivpusic/grpool"
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
	fName := "dcp-sim-error-log-" + strconv.Itoa(int(time.Now().UnixNano())) + ".txt"
	fp := filepath.Join("log", fName)
	f, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		panic(err.Error())
	}

	defer f.Close()

	redirectStderr(f)

	if len(os.Args) < 3 {
		fmt.Println("Program require an arg in set [1, " + strconv.Itoa(NrOfAvailableRuns) + "] and a duration in seconds")
		return
	}

	runNr, e := strconv.Atoi(os.Args[1])
	if e != nil {
		fmt.Println("Run nr not an integer")
		return
	}

	runDuration, e := strconv.Atoi(os.Args[2])
	if e != nil {
		fmt.Println("Timer duration not an integer")
		return
	}

	if runNr < 1 || runNr > NrOfAvailableRuns {
		fmt.Println("Argument must be in set [1, " + strconv.Itoa(NrOfAvailableRuns) + "]")
		return
	}

	//os.Stdout = os.Stderr
	os.Stdout = nil

	rc := RunCaller{}
	runConfig := rc.GetRunConfig(runNr)

	startTime := time.Now()
	runSimulation(runConfig, startTime, runDuration)
}

func redirectStderr(f *os.File) {
	err := syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		log.Fatalf("Failed to redirect stderr to file: %v", err)
	}
}

func createNodes(numberOfNodes int, config DCP.CtNodeConfig, poolPtr *grpool.Pool) []*DCP.CtNode {
	var nodes []*DCP.CtNode
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < numberOfNodes; i++ {
		node := DCP.NewCtNode(GenerateIdTable(rand.Intn(25-1)+1), config, poolPtr)

		e := node.Co.KeyGen()
		if e != nil {
			fmt.Println(e.Error())
			break
		}

		nodes = append(nodes, node)
	}

	return nodes
}

func runSimulation(runConfig RunConfig, startTime time.Time, runTime int) {
	config := DCP.NewCtNodeConfig()
	config.NodeVisitDecryptThreshold = runConfig.DecryptThreshold
	config.Throttle = &runConfig.Latency
	config.CoTTL = runConfig.TTL

	// Non-run specific config
	config.IncludeHistory = false
	config.SuppressLogging = false

	pool := grpool.NewPool(10000, 100)
	defer pool.Release()

	nodes := createNodes(runConfig.NrOfNodes, config, pool)

	if runConfig.Topology == Cluster {
		EstablishNodeRelationshipsLocalClusters(nodes, *runConfig.ClusterSize)
	} else {
		EstablishNodeRelationShipAllInRange(nodes)
	}

	lock1 := &sync.RWMutex{}

	closeMonitor := make(chan struct{})
	go LaunchMonitor(nodes, closeMonitor, lock1, true)

	for _, node := range nodes {
		node.Listen()

		go func(node *DCP.CtNode) {
			CalculationProcessInitiator(node)
		}(node)
	}

	for {
		select {
		case <-closeMonitor:
			e := generateReport(nodes, runConfig)
			if e != nil {
				panic(e)
			}

			return
		default:
			timeNow := time.Now()

			if timeNow.Sub(startTime).Seconds() > float64(runTime) {
				close(closeMonitor)
			}

			nGoR := runtime.NumGoroutine()
			nBranches := DCP.GetNrOfActiveBranches()
			fmt.Println(nGoR, nBranches)
		}
	}
}
