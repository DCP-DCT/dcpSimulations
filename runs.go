package main

const NrOfAvailableRuns = 25

type RunCaller struct{}

var ClusterSize = 8

func (rc RunCaller) GetRunConfig(nr int) RunConfig {
	switch nr {
	case 1:
		return Run1()
	case 2:
		return Run2()
	case 3:
		return Run3()
	case 4:
		return Run4()
	case 5:
		return Run5()
	case 6:
		return Run6()
	case 7:
		return Run7()
	case 8:
		return Run8()
	case 9:
		return Run9()
	case 10:
		return Run10()
	case 24:
		return Run24()
	case 25:
		return Run25()
	case 11:
		return Run11()
	case 12:
		return Run12()
	case 13:
		return Run13()
	case 23:
		return Run23()
	case 14:
		return Run14()
	case 15:
		return Run15()
	case 16:
		return Run16()
	case 17:
		return Run17()
	case 18:
		return Run18()
	case 19:
		return Run19()
	case 20:
		return Run20()
	case 21:
		return Run21()
	case 22:
		return Run22()
	default:
		return RunConfig{}
	}
}

// #####START PRIVACY RUNS#####
func Run1() RunConfig {
	return RunConfig{
		RunDescription:   "1",
		NrOfNodes:        100,
		Latency:          30,
		DecryptThreshold: 3,
		TTL:              6,
		Topology:         Cluster,
		ClusterSize:      &ClusterSize,
		DropAfterNrDuplicates: 1,
	}
}

func Run2() RunConfig {
	return RunConfig{
		RunDescription:   "2",
		NrOfNodes:        100,
		Latency:          30,
		DecryptThreshold: 4,
		TTL:              6,
		Topology:         Cluster,
		ClusterSize:      &ClusterSize,
		DropAfterNrDuplicates: 2,
	}
}

func Run3() RunConfig {
	return RunConfig{
		RunDescription:   "3",
		NrOfNodes:        100,
		Latency:          30,
		DecryptThreshold: 5,
		TTL:              9,
		Topology:         Cluster,
		ClusterSize:      &ClusterSize,
		DropAfterNrDuplicates: 4,
	}
}

func Run4() RunConfig {
	return RunConfig{
		RunDescription:   "4",
		NrOfNodes:        100,
		Latency:          30,
		DecryptThreshold: 6,
		TTL:              60,
		Topology:         Cluster,
		ClusterSize:      &ClusterSize,
		DropAfterNrDuplicates: 60,
	}
}

func Run5() RunConfig {
	return RunConfig{
		RunDescription:   "5",
		NrOfNodes:        25,
		Latency:          30,
		DecryptThreshold: 9,
		TTL:              10,
		Topology:         All,
		ClusterSize:      nil,
		DropAfterNrDuplicates: -1,
	}
}

func Run6() RunConfig {
	return RunConfig{
		RunDescription:   "6",
		NrOfNodes:        100,
		Latency:          30,
		DecryptThreshold: 9,
		TTL:              10,
		Topology:         All,
		ClusterSize:      nil,
		DropAfterNrDuplicates: -1,
	}
}

// #####END PRIVACY RUNS#####

// #####EFFICIENCY RUNS#####
func Run7() RunConfig {
	cs := 8
	return RunConfig{
		RunDescription:   "7",
		NrOfNodes:        100,
		Latency:          30,
		DecryptThreshold: 3,
		TTL:              3,
		Topology:         Cluster,
		ClusterSize:      &cs,
		DropAfterNrDuplicates: -1,
	}
}

func Run8() RunConfig {
	cs := 8
	return RunConfig{
		RunDescription:   "8",
		NrOfNodes:        100,
		Latency:          30,
		DecryptThreshold: 3,
		TTL:              6,
		Topology:         Cluster,
		ClusterSize:      &cs,
		DropAfterNrDuplicates: -1,
	}
}

func Run9() RunConfig {
	cs := 8
	return RunConfig{
		RunDescription:   "9",
		NrOfNodes:        100,
		Latency:          30,
		DecryptThreshold: 3,
		TTL:              9,
		Topology:         Cluster,
		ClusterSize:      &cs,
		DropAfterNrDuplicates: -1,
	}
}

func Run10() RunConfig {
	cs := 8
	return RunConfig{
		RunDescription:   "10",
		NrOfNodes:        100,
		Latency:          30,
		DecryptThreshold: 3,
		TTL:              12,
		Topology:         Cluster,
		ClusterSize:      &cs,
		DropAfterNrDuplicates: -1,
	}
}

func Run24() RunConfig {
	cs := 8
	return RunConfig{
		RunDescription:   "24",
		NrOfNodes:        100,
		Latency:          30,
		DecryptThreshold: 3,
		TTL:              15,
		Topology:         Cluster,
		ClusterSize:      &cs,
		DropAfterNrDuplicates: -1,
	}
}

func Run25() RunConfig {
	cs := 8
	return RunConfig{
		RunDescription:   "25",
		NrOfNodes:        100,
		Latency:          30,
		DecryptThreshold: 3,
		TTL:              18,
		Topology:         Cluster,
		ClusterSize:      &cs,
		DropAfterNrDuplicates: -1,
	}
}

// #####END EFFICIENCY RUNS#####

// #####COMPLEXITY RUNS#####
func Run11() RunConfig {
	return RunConfig{
		RunDescription:   "11",
		NrOfNodes:        25,
		Latency:          30,
		DecryptThreshold: 3,
		TTL:              10,
		Topology:         All,
		ClusterSize:      nil,
		DropAfterNrDuplicates: -1,
	}
}

func Run12() RunConfig {
	return RunConfig{
		RunDescription:   "12",
		NrOfNodes:        100,
		Latency:          30,
		DecryptThreshold: 3,
		TTL:              10,
		Topology:         All,
		ClusterSize:      nil,
		DropAfterNrDuplicates: -1,
	}
}

func Run13() RunConfig {
	return RunConfig{
		RunDescription:   "13",
		NrOfNodes:        200,
		Latency:          30,
		DecryptThreshold: 3,
		TTL:              10,
		Topology:         All,
		ClusterSize:      nil,
		DropAfterNrDuplicates: -1,
	}
}

func Run23() RunConfig {
	return RunConfig{
		RunDescription:   "23",
		NrOfNodes:        300,
		Latency:          30,
		DecryptThreshold: 3,
		TTL:              10,
		Topology:         All,
		ClusterSize:      nil,
		DropAfterNrDuplicates: -1,
	}
}

func Run14() RunConfig {
	return RunConfig{
		RunDescription:   "14",
		NrOfNodes:        600,
		Latency:          30,
		DecryptThreshold: 3,
		TTL:              10,
		Topology:         All,
		ClusterSize:      nil,
		DropAfterNrDuplicates: -1,
	}
}

func Run15() RunConfig {
	cs := 8
	return RunConfig{
		RunDescription:   "15",
		NrOfNodes:        25,
		Latency:          30,
		DecryptThreshold: 3,
		TTL:              10,
		Topology:         Cluster,
		ClusterSize:      &cs,
		DropAfterNrDuplicates: -1,
	}
}

func Run16() RunConfig {
	cs := 8
	return RunConfig{
		RunDescription:   "16",
		NrOfNodes:        100,
		Latency:          30,
		DecryptThreshold: 3,
		TTL:              10,
		Topology:         Cluster,
		ClusterSize:      &cs,
		DropAfterNrDuplicates: -1,
	}
}

func Run17() RunConfig {
	cs := 8
	return RunConfig{
		RunDescription:   "17",
		NrOfNodes:        200,
		Latency:          30,
		DecryptThreshold: 3,
		TTL:              10,
		Topology:         Cluster,
		ClusterSize:      &cs,
		DropAfterNrDuplicates: -1,
	}
}

func Run18() RunConfig {
	cs := 8
	return RunConfig{
		RunDescription:   "18",
		NrOfNodes:        600,
		Latency:          30,
		DecryptThreshold: 3,
		TTL:              10,
		Topology:         Cluster,
		ClusterSize:      &cs,
		DropAfterNrDuplicates: -1,
	}
}

func Run19() RunConfig {
	return RunConfig{
		RunDescription:   "19",
		NrOfNodes:        100,
		Latency:          100,
		DecryptThreshold: 3,
		TTL:              10,
		Topology:         All,
		ClusterSize:      nil,
		DropAfterNrDuplicates: -1,
	}
}

func Run20() RunConfig {
	return RunConfig{
		RunDescription:   "20",
		NrOfNodes:        100,
		Latency:          200,
		DecryptThreshold: 3,
		TTL:              10,
		Topology:         All,
		ClusterSize:      nil,
		DropAfterNrDuplicates: -1,
	}
}

func Run21() RunConfig {
	cs := 8
	return RunConfig{
		RunDescription:   "21",
		NrOfNodes:        100,
		Latency:          100,
		DecryptThreshold: 3,
		TTL:              10,
		Topology:         Cluster,
		ClusterSize:      &cs,
		DropAfterNrDuplicates: -1,
	}
}

func Run22() RunConfig {
	cs := 8
	return RunConfig{
		RunDescription:   "22",
		NrOfNodes:        100,
		Latency:          200,
		DecryptThreshold: 3,
		TTL:              10,
		Topology:         Cluster,
		ClusterSize:      &cs,
		DropAfterNrDuplicates: -1,
	}
}

// #####ENDCOMPLEXITY RUNS#####