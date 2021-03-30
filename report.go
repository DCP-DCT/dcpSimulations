package main

import (
	"encoding/json"
	"github.com/DCP-DCT/DCP"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"time"
)

func generateReport(nodes []*DCP.CtNode) error {
	fName := "dcp-sim-report-" + strconv.Itoa(int(time.Now().UnixNano())) + ".json"

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
