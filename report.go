package main

import (
	"encoding/json"
	"github.com/DCP-DCT/DCP"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"time"
)

func generateReport(nodes []*DCP.CtNode, info RunConfig) error {

	fName := "dcp-report-" + info.RunDescription + "-" + strconv.Itoa(int(time.Now().UnixNano())) + ".json"

	// Program riddled with read/writes races so just wait and hope they all finish before processing maps
	time.Sleep(20 * time.Second)

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
