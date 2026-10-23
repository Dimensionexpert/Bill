package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/Dimensionexpert/Bills/internal/helpers"
)

func main() {
	file, err := os.Open("./source/s68.csv")
	if err != nil {
		fmt.Println("can't open file:", err)
		return
	}
	defer file.Close()

	fmt.Println(file.Name())

	record, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("can't read CSV:", err)
		return
	}

	blockidx, err := helpers.GetBlockInfoIDs(record)
	if err != nil {
		fmt.Println("can't get block indices:", err)
		return
	}

	blockNamenCode, err := helpers.GetBlockInfo(record, blockidx)
	if err != nil {
		fmt.Println("can't get the block info:", err)
	}

	clusteridx, err := helpers.GetClusterInfoIDs(record)
	if err != nil {
		fmt.Println("can't get block indices:", err)
		return
	}

	clusterNamenCode, err := helpers.GetClusterInfo(record, clusteridx)
	if err != nil {
		fmt.Println("can't get the block info:", err)
	}

	fmt.Println(blockNamenCode)
	fmt.Println("lenght of cluster idx", clusteridx)
	fmt.Println("length of cluster retrived", clusterNamenCode)

}
