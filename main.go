package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/Dimensionexpert/Bills/internal/helpers"
)

func main() {
	file, err := os.Open("./source/s15.csv")
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

	clusterCodeNameidx, err := helpers.GetClusterCodeNameIDs(record)
	if err != nil {
		fmt.Println("can't get block indices:", err)
		return
	}

	clusterNamenCode, err := helpers.GetClusterInfo(record, clusterCodeNameidx)
	if err != nil {
		fmt.Println("can't get the block info:", err)
	}

	clusterSeveredAmountidx, err := helpers.GetClusterServedAmountIDs(record)
	if err != nil {
		fmt.Println("can't get the clusterTotal id info:", err)
	}

	clusterSeveredAmount, err := helpers.GetClusterServedAmount(record, clusterSeveredAmountidx)
	if err != nil {
		fmt.Println("can't get the clusterTotal info:", err)
	}

	blockSeveredAmountidx, err := helpers.GetBlockServedAmountIDs(record)
	if err != nil {
		fmt.Println("can't get the clusterTotal id info:", err)
	}
	blockSeveredAmount, err := helpers.GetBlockServedAmount(record,blockSeveredAmountidx)
	if err != nil {
		fmt.Println("can't get the clusterTotal id info:", err)
	}

	fmt.Println(blockidx)       // clean and expected output
	fmt.Println(blockNamenCode) // clean and expected output

	fmt.Println(clusterCodeNameidx) // clean and expected output
	fmt.Println(clusterNamenCode)   // clean and expected output

	fmt.Println(clusterSeveredAmountidx) // clean and expected output
	fmt.Println(clusterSeveredAmount)    // clean and expected output

	fmt.Println(blockSeveredAmountidx) // clean and expected output
	fmt.Println(blockSeveredAmount) // clean and expected output
	

}
