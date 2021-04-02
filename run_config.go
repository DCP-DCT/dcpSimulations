package main

import "time"

type Topology int

const (
	Cluster Topology = iota
	All
)

type RunConfig struct {
	RunDescription   string
	NrOfNodes        int
	Latency          time.Duration
	DecryptThreshold int
	TTL              int
	Topology         Topology
	ClusterSize      *int
}
