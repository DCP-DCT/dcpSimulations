package main

import (
	"fmt"
	"github.com/DCP-DCT/DCP"
	"math/rand"
	"testing"
	"time"
)

func TestEstablishNodeRelationshipsLocalClusters(t *testing.T) {
	var nodes []*DCP.CtNode
	rand.Seed(time.Now().UnixNano())

	config := DCP.NewCtNodeConfig()

	for i := 0; i < 10; i++ {
		node := DCP.NewCtNode([]string{"foo"}, config)

		e := node.Co.KeyGen()
		if e != nil {
			fmt.Println(e.Error())
			break
		}

		nodes = append(nodes, node)
	}

	cs := 3
	csP := &cs
	EstablishNodeRelationshipsLocalClusters(nodes, *csP)

	if _, exist := nodes[2].TransportLayer.ReachableNodes[nodes[3].TransportLayer.DataCh]; !exist {
		t.Fail()
	}
}
