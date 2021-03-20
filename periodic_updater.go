package main

import (
	"fmt"
	"github.com/DCP-DCT/DCP"
	"math/rand"
	"time"
)

// randomCalculationProcessInitiator takes a single node and based
// on a random interval will begin an broadcast process.
//
// Will select a number at random {1, 10} and if the number is X then
// the process will be initiated or else sleep for some period.
func RandomCalculationProcessInitiator(node *DCP.CtNode, stop chan struct{}) {
	rand.Seed(time.Now().UnixNano())

	for {
		select {
			case <- stop:
				fmt.Println("Stop received, shutting down RandomCalculationProcessInitiator")
				break
		default:
			nr := rand.Intn(10 - 1) + 1

			if nr == 5 {
				e := DCP.InitRoutine(DCP.PrepareIdLenCalculation, node)
				if e != nil {
					fmt.Println(e)
					return
				}

				fmt.Printf("Starting process for node %s\n", node.Id)
				node.Broadcast(nil)
				return
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func RandomReachableNodeShuffler(node *DCP.CtNode, stop chan struct{}) {}