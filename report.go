package main

import (
	"encoding/json"
	"github.com/DCP-DCT/DCP"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"time"
)

func generateReport(nodes []*DCP.CtNode, cluster bool) error {
	topologyIndicator := ""
	if cluster {
		topologyIndicator = "cluster-"
	}

	fName := "dcp-sim-report-" + topologyIndicator + strconv.Itoa(int(time.Now().UnixNano())) + ".json"

	// Program riddled with race read/writes so just wait and hope they all finish before marshal
	time.Sleep(500 * time.Millisecond)

	b, e := json.Marshal(nodes)
	if e != nil {
		return e
	}

	fp := filepath.Join("reports", fName)

	e = ioutil.WriteFile(fp, b, 0644)
	if e != nil {
		return e
	}

	return nil
}
