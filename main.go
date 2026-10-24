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
	blockSeveredAmount, err := helpers.GetBlockServedAmount(record, blockSeveredAmountidx)
	if err != nil {
		fmt.Println("can't get the clusterTotal id info:", err)
	}

	schoolIdx, err := helpers.GetSchoolInfoIdx(record, blockNamenCode[0].BlockCode)
	if err != nil {
		fmt.Println("can't get the school info indices info:", err)
	}

	schoolInfo, err := helpers.GetUDISENameServedAmount(record, schoolIdx)
	if err != nil {
		fmt.Println("can't get the school info :", err)
	}

	fmt.Printf("Block Name:%s  Code:%s\n", blockNamenCode[0].BlockName, blockNamenCode[0].BlockCode)                                                                             // checked ! clean and expected output
	fmt.Printf("Block Served:%s  Total:%s\n", blockSeveredAmount[0].Served, blockSeveredAmount[0].Amount)                                                                        // checked ! clean and expected output
	fmt.Printf("Cluster Name:%s  Code:%s\n", clusterNamenCode[0].ClusterName, clusterNamenCode[0].ClusterCode)                                                                   // checked ! clean and expected output
	fmt.Printf("Cluster Served:%s  Total:%s\n", clusterSeveredAmount[0].Served, clusterSeveredAmount[0].Amount)                                                                  // checked ! Not expected output
	fmt.Printf("School UDISE : %s Name : %s Meal served : %s Amount : %s\n", schoolInfo[0].UDISECode, schoolInfo[0].SchoolName, schoolInfo[0].MealsServed, schoolInfo[0].Amount) // checked ! clean and expected output

}
