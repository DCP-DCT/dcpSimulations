package main

import (
	"fmt"
	"github.com/DCP-DCT/DCP"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/olekukonko/tablewriter"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func LaunchMonitor(nodes []*DCP.CtNode, done chan struct{}) {
	if err := ui.Init(); err != nil {
		log.Fatal("Could not initialize monitor")
	}
	defer ui.Close()

	nodeList := widgets.NewList()
	info := createInfoParagraph(0, len(nodes))
	createList(nodeList, nodes)
	ui.Render(nodeList)

	ticker := time.NewTicker(time.Second).C
	tickCount := 0
	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			{
				switch e.ID {
				case "q", "<C-c>":
					fmt.Println()
					close(done)
					return
				case "j", "<Down>":
					nodeList.ScrollDown()
					ui.Render(info, nodeList)
				case "k", "<Up>":
					nodeList.ScrollUp()
					ui.Render(info, nodeList)
				}
			}
		case <-ticker:
			info = createInfoParagraph(tickCount, len(nodes))
			createList(nodeList, nodes)

			ui.Render(info, nodeList)
			tickCount++
		}
	}
}

/*func createNodeDisplayListItemRow(node *DCP.CtNode) string {
	return
}*/

func createNodeDisplayListItemTable(nodes []*DCP.CtNode) string {
	var data [][]string

	for _, node := range nodes {
		cip := big.NewInt(0)
		if node.Co.Cipher != nil {
			cip = node.Co.Decrypt(node.Co.Cipher)
		}

		data = append(data, []string{
			node.Id.String(),
			strconv.Itoa(node.Co.Counter),
			cip.String(),
			strconv.FormatBool(node.IsCalculationProcessRunning()),
			strconv.Itoa(node.Diagnosis.NumberOfUpdates),
			strconv.Itoa(node.Diagnosis.NumberOfRejectedDueToThreshold),
			strconv.Itoa(node.Diagnosis.NumberOfBroadcasts),
			strconv.Itoa(node.Diagnosis.NumberOfDuplicates),
			strconv.Itoa(node.Diagnosis.NumberOfPkMatches),
			strconv.Itoa(node.Diagnosis.NumberOfPacketsDropped),
		})
	}

	header := []string{
		"nodeId",
		"counter",
		"cipher decr",
		"Calc running",
		"Updates",
		"Rejects",
		"Broadcasts",
		"Duplicates",
		"Pk matches",
		"Expired TTLs",
	}

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader(header)
	table.AppendBulk(data)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetTablePadding("\t")
	table.Render()
	return tableString.String()
}

func createList(list *widgets.List, nodes []*DCP.CtNode) {
	listItems := createNodeDisplayListItemTable(nodes)

	w, _ := ui.TerminalDimensions()

	list.Title = "Nodes"
	list.Rows = []string{listItems}
	list.WrapText = true
	list.Border = false
	list.SetRect(0, 10, w, 40)
}

func createInfoParagraph(tickCount int, nrOfNodes int) *widgets.Paragraph {
	w, _ := ui.TerminalDimensions()

	info := widgets.NewParagraph()
	info.Title = "Run information"
	info.Text = "Tick count: " + strconv.Itoa(tickCount) + "\nNode count: " + strconv.Itoa(nrOfNodes)
	info.SetRect(0, 0, w, 10)

	return info
}
