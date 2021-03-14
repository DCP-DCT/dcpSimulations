package main

import (
	"github.com/DCP-DCT/DCP"
	"testing"
)

func TestContains (t *testing.T) {
	chan1 := make(chan *DCP.CalculationObjectPaillier)
	chan2 := make(chan *DCP.CalculationObjectPaillier)
	chan3 := make(chan *DCP.CalculationObjectPaillier)

	defer close(chan1)
	defer close(chan2)
	defer close(chan3)

	source := []chan*DCP.CalculationObjectPaillier{chan1, chan2}

	if !contains(source, chan1) {
		t.Fail()
	}

	if contains(source, chan3) {
		t.Fail()
	}
}