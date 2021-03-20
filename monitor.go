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

func LaunchMonitor(nodes []*DCP.CtNode, done chan struct{}) {
	if err := ui.Init(); err != nil {
		log.Fatal("Could not initialize monitor")
	}
	defer ui.Close()

	grid := ui.NewGrid()
	w, h := ui.TerminalDimensions()
	grid.SetRect(0, 0, w, h)

	rows := createGridEntries(nodes, 0)

	grid.Set(rows...)

	ui.Render(grid)

	uiEvents := ui.PollEvents()
	ticker := time.NewTicker(time.Second).C
	tickCount := 0
	for {
		select {
		case e := <-uiEvents:
			{
				switch e.ID {
				case "q", "<C-c>":
					close(done)
					return
				}
			}
		case <-ticker:
			rows = createGridEntries(nodes, tickCount)
			grid.Set(rows...)
			ui.Render(grid)
			tickCount++
		}
	}
}

func createNodeDisplay(node *DCP.CtNode) *widgets.Table {
	table := widgets.NewTable()
	table.Rows = [][]string{
		{"nodeId", "counter", "Calculation process running"},
		{node.Id.String(), strconv.Itoa(node.Co.Counter), strconv.FormatBool(node.IsCalculationProcessRunning())},
	}

	if table == nil {
		fmt.Println("Table nil")
	}
	return table
}

func createGridEntries(nodes []*DCP.CtNode, tickCount int) []interface{} {
	ratioCoefficient := 2
	colRatio := 1.0 / float64(ratioCoefficient)
	rowRatio := 1.0 / (float64(len(nodes)) + 1/float64(ratioCoefficient))

	var cols []interface{}
	for _, node := range nodes {
		cols = append(cols, ui.NewCol(colRatio, createNodeDisplay(node)))
	}

	var rows []interface{}

	p := widgets.NewParagraph()
	p.Text = "Tick count: " + strconv.Itoa(tickCount)
	rows = append(rows, ui.NewRow(rowRatio, ui.NewCol(1.0, p)))

	it := 0
	for {
		if len(cols) == 0 {
			break
		}

		if it+ratioCoefficient > len(cols)-1 {
			currentCols := cols[it:]
			rows = append(rows, ui.NewRow(rowRatio, currentCols...))
			break
		}

		currentCols := cols[it : it+ratioCoefficient]

		rows = append(rows, ui.NewRow(rowRatio, currentCols...))

		it = it + ratioCoefficient
	}

	return rows
}
