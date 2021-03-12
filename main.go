package main

import (
	"fmt"
	"github.com/DCP-DCT/DCP"
)

func main() {
	nodes := make([]DCP.CtNode, 10)

	for _, node := range nodes {
		node.Co.KeyGen()
		node.Ids = GenerateIdTable(5)
		fmt.Println(node)
	}
}
