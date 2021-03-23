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

	w, _ := ui.TerminalDimensions()

	info := createInfoParagraph(0, len(nodes))

	nodeList := widgets.NewList()
	createList(nodeList, nodes)

	actionsList := widgets.NewList()
	createActionsList(actionsList, nodes)

	tp := widgets.NewTabPane("Runtime data", "Control actions")
	tp.SetRect(0, 10, w, 13)
	tp.Border = true

	scrollList := nodeList

	renderTab := func() {
		switch tp.ActiveTabIndex {
		case 0:
			scrollList = nodeList
			ui.Render(nodeList)
		case 1:
			scrollList = actionsList
			ui.Render(actionsList)
		}
	}

	ui.Render(info, tp)

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
					scrollList.ScrollDown()
					ui.Render(info, tp)
				case "k", "<Up>":
					scrollList.ScrollUp()
					ui.Render(info, tp)
				case "h":
					tp.FocusLeft()
					ui.Clear()
					ui.Render(info, tp)
					renderTab()
				case "l":
					tp.FocusRight()
					ui.Clear()
					ui.Render(info, tp)
					renderTab()
				}
			}
		case <-ticker:
			info = createInfoParagraph(tickCount, len(nodes))
			createList(nodeList, nodes)
			createActionsList(actionsList, nodes)
			renderTab()

			ui.Render(info, tp)
			tickCount++
		}
	}
}

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

	list.Rows = []string{listItems}
	list.WrapText = true
	list.Border = true

	w, _ := ui.TerminalDimensions()

	list.SetRect(0, 13, w, 40)
}

func createActionsList(list *widgets.List, nodes []*DCP.CtNode) {
	var allRecords []ContributionRecord
	for _, node := range nodes {
		allRecords = append(allRecords, ContributionRecord{
			NodeId:  node.Id,
			Updates: node.Diagnosis.Control.NodesContributedToUpdates,
		})
	}

	var listItems []string
	for _, record := range allRecords {
		for co, contribution := range record.Updates {
			listItems = append(listItems, "NodeId: "+record.NodeId.String()+" -> Id/BranchId ["+co.Id.String()+"/"+co.BranchId.String()+"] Added: "+strconv.Itoa(contribution))
		}
	}

	list.Rows = listItems
	list.WrapText = true
	list.Border = true
	list.SelectedRowStyle.Fg = ui.ColorMagenta

	w, _ := ui.TerminalDimensions()

	list.SetRect(0, 13, w, 40)
}

func createInfoParagraph(tickCount int, nrOfNodes int) *widgets.Paragraph {
	w, _ := ui.TerminalDimensions()

	info := widgets.NewParagraph()
	info.Title = "Run information"
	info.Text = "Tick count: " + strconv.Itoa(tickCount) + "\nNode count: " + strconv.Itoa(nrOfNodes)
	info.SetRect(0, 0, w, 10)

	return info
}
