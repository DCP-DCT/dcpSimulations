package main

import (
	"fmt"
	"github.com/DCP-DCT/DCP"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"log"
	"strconv"
	"time"
)


func LaunchMonitor(nodes []*DCP.CtNode) {
	if err := ui.Init(); err != nil {
		log.Fatal("Could not initialize monitor")
	}
	defer ui.Close()

	grid := ui.NewGrid()
	w, h := ui.TerminalDimensions()
	grid.SetRect(0, 0, w, h)

	rows := createGridEntries(nodes)

	grid.Set(rows...)

	ui.Render(grid)

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	for {
		select {
		case e := <-uiEvents: {
				switch e.ID {
				case "q", "<C-c>":
					return
				}
			}
			case <- ticker:
				rows = createGridEntries(nodes)
				grid.Set(rows...)
				ui.Render(grid)
		}
	}
}

func createNodeDisplay(node *DCP.CtNode) *widgets.Table {
	table := widgets.NewTable()
	table.Rows = [][]string{
		{"nodeId", "counter"},
		{node.Id.String(), strconv.Itoa(node.Co.Counter)},
	}

	if table == nil {
		fmt.Println("Table nil")
	}
	return table
}

func createGridEntries(nodes []*DCP.CtNode) []interface{} {
	ratioCoefficient := 4
	colRatio := 1.0 / float64(ratioCoefficient)
	rowRatio := 1.0 / (float64(len(nodes)) / float64(ratioCoefficient))

	var cols []interface{}
	for _, node := range nodes {
		cols = append(cols, ui.NewCol(colRatio, createNodeDisplay(node)))
	}

	var rows []interface{}
	it := 0
	for {
		if len(cols) == 0 {
			break
		}

		if it + ratioCoefficient > len(cols) -1 {
			currentCols := cols[it:]
			rows = append(rows, ui.NewRow(rowRatio, currentCols...))
			break
		}

		currentCols := cols[it:it+ratioCoefficient]

		rows = append(rows, ui.NewRow(rowRatio, currentCols...))

		it = it + ratioCoefficient
	}

	return rows
}
