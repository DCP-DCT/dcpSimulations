package main

import (
	"github.com/DCP-DCT/DCP"
	"math/rand"
	"time"
)

// EstablishNodeRelationships takes an array of CtNodes and
// randomly sets the CtNode.ReachableNodes attribute on the
// supplied nodes.
func EstablishNodeRelationships(nodes []*DCP.CtNode, initialNode *DCP.CtNode) {
	if len(nodes) < 2 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(nodes); i++ {
		current := nodes[i]

		numbersToAdd := rand.Intn(len(nodes)-1) + 1

		for j := 0; j < numbersToAdd; j++ {
			// randomNodeIndex between {0, len(nodes)}
			var randomNodeIndex int
			it := 1
			for {
				it++
				randomNodeIndex = rand.Intn(len(nodes)-1) + 1

				if _, exists := current.TransportLayer.ReachableNodes[nodes[randomNodeIndex].TransportLayer.DataCh]; exists {
					break
				}

				if current.Id != nodes[randomNodeIndex].Id {
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

			current.TransportLayer.ReachableNodes[reachableNode.TransportLayer.DataCh] = struct{}{}
		}
	}

	// ensure last node has a connection back to initial node
	if _, exists := nodes[len(nodes)-1].TransportLayer.ReachableNodes[initialNode.TransportLayer.DataCh]; !exists {
		nodes[len(nodes)-1].TransportLayer.ReachableNodes[initialNode.TransportLayer.DataCh] = struct{}{}
	}
}

func EstablishNodeRelationShipAllInRange(nodes []*DCP.CtNode) {
	allTransportLayers := make(map[chan []byte]struct{})
	for _, node := range nodes {
		allTransportLayers[node.TransportLayer.DataCh] = struct{}{}
	}

	for _, node := range nodes {
		for k, v := range allTransportLayers {
			if k != node.TransportLayer.DataCh {
				node.TransportLayer.ReachableNodes[k] = v
			}
		}
	}
}

func EstablishNodeRelationshipsLocalClusters(nodes []*DCP.CtNode, maxSizeCluster int) {
	var clusters [][]*DCP.CtNode
	offset := 0
	for {
		if (offset + maxSizeCluster) < len(nodes) {
			cluster := nodes[offset : offset+maxSizeCluster]
			clusters = append(clusters, cluster)
			offset = offset + maxSizeCluster
		} else {
			cluster := nodes[offset:]
			clusters = append(clusters, cluster)
			break
		}
	}

	for i, cluster := range clusters {
		clusterTransportLayers := make(map[chan []byte]struct{})
		for _, node := range cluster {
			clusterTransportLayers[node.TransportLayer.DataCh] = struct{}{}
		}

		for j, node := range cluster {
			for k, v := range clusterTransportLayers {
				if k != node.TransportLayer.DataCh {
					node.TransportLayer.ReachableNodes[k] = v
				}
			}

			// Last node in cluster, assign link to first node in next cluster
			if j == len(cluster)-1 {
				if i+1 <= len(clusters)-1 {
					firstNodeNextCluster := clusters[i+1][0]
					node.TransportLayer.ReachableNodes[firstNodeNextCluster.TransportLayer.DataCh] = struct{}{}
				}
			}
		}
	}
}

// Mesh topology

// Free create, more tails
