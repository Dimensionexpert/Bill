package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/Dimensionexpert/Bills/internal/database"
	"github.com/Dimensionexpert/Bills/internal/helpers"
	"github.com/Dimensionexpert/Bills/internal/validate"
)

func main() {
	// ---------------------------------
	// open the csv
	// ---------------------------------
	file, err := os.Open("./source/Jun/c68.csv")
	if err != nil {
		fmt.Println(":( can't open file:", err)
		return
	}
	defer file.Close()

	fmt.Println(file.Name())

	record, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(":( can't read CSV:", err)
		return
	}
	// ---------------------------------
	// Parse block information
	// ---------------------------------
	blockidx, err := helpers.GetBlockInfoIDs(record)
	if err != nil {
		fmt.Println(":( can't get block indices:", err)
		return
	}

	blockNamenCode, err := helpers.GetBlockInfo(record, blockidx)
	if err != nil {
		fmt.Println(":( can't get the block info:", err)
	}

	// ---------------------------------
	// Parse cluster information
	// ---------------------------------
	clusterCodeNameidx, err := helpers.GetClusterCodeNameIDs(record)
	if err != nil {
		fmt.Println(":( can't get cluster indices:", err)
		return
	}

	clusterNamenCode, err := helpers.GetClusterInfo(record, clusterCodeNameidx)
	if err != nil {
		fmt.Println(":( can't get the cluster info:", err)
	}

	// ---------------------------------
	// Parse cluster totals
	// ---------------------------------
	clusterSeveredAmountidx, err := helpers.GetClusterServedAmountIDs(record)
	if err != nil {
		fmt.Println(":( can't get the cluster total IDs:", err)
	}

	clusterSeveredAmount, err := helpers.GetClusterServedAmount(record, clusterSeveredAmountidx)
	if err != nil {
		fmt.Println(":( can't get the cluster totals:", err)
	}

	// ---------------------------------
	// Parse block totals
	// ---------------------------------
	blockSeveredAmountidx, err := helpers.GetBlockServedAmountIDs(record)
	if err != nil {
		fmt.Println(":( can't get the block total IDs:", err)
	}

	blockSeveredAmount, err := helpers.GetBlockServedAmount(record, blockSeveredAmountidx)
	if err != nil {
		fmt.Println(":( can't get the block totals:", err)
	}

	// ---------------------------------
	// Parse school information
	// ---------------------------------
	schoolIdx, err := helpers.GetSchoolInfoIdx(record, blockNamenCode[0].BlockCode)
	if err != nil {
		fmt.Println(":( can't get the school indices:", err)
	}

	schoolInfo, err := helpers.GetSchools(record, schoolIdx)
	if err != nil {
		fmt.Println(":( can't get the school information:", err)
	}

	// ---------------------------------
	// Group schools by cluster
	// ---------------------------------
	schoolByClusteridx, err := helpers.SchoolsByClusterIdx(
		clusterCodeNameidx,
		clusterSeveredAmountidx,
		schoolIdx,
	)
	if err != nil {
		fmt.Println(":( can't get school indices by cluster:", err)
	}

	schoolsByclusterInfo, err := helpers.SchoolsByClusterinfo(
		record,
		clusterNamenCode,
		schoolByClusteridx,
	)
	if err != nil {
		fmt.Println(":( can't group schools by cluster:", err)
	}

	// ---------------------------------
	// validation layer
	// ---------------------------------

	// ---------------------------------
	// cluster wise totals
	// ---------------------------------
	clusterWiseServedTotal, clusterWiseAmountTotal, err := validate.GetTotalAmountsCluster(schoolsByclusterInfo)
	if err != nil {
		fmt.Printf(":( calculating cluster totals failed: %v\n", err)
	}

	clusterOk, err := validate.ValidateClusterTotals(
		clusterSeveredAmount,
		clusterWiseServedTotal,
		clusterWiseAmountTotal,
	)

	if err != nil {
		fmt.Printf(":( Cluster validation failed: %v\n", err)
		return
	}

	if !clusterOk {
		fmt.Println(":( Cluster validation failed")
		return
	}

	fmt.Println(":) Cluster validation passed")

	// ---------------------------------
	// block wise validation
	// ---------------------------------

	blockWiseServedTotal, blockWiseAmountTotal, err := validate.GetTotalAmountsBlock(clusterSeveredAmount)
	if err != nil {
		fmt.Printf(":( calculating block totals failed: %v\n", err)
	}

	blockOk, err := validate.ValidateBlockTotals(blockSeveredAmount, blockWiseServedTotal, blockWiseAmountTotal)
	if err != nil {
		fmt.Printf(":( Block validation failed: %v\n", err)
		return
	}

	if !blockOk {
		fmt.Println(":( Block validation failed")
		return
	}

	fmt.Println(":) Block validation passed")

	// ---------------------------------
	// data cross checking
	// ---------------------------------

	fmt.Printf("Block Name:%s  Code:%s\n", blockNamenCode[0].BlockName, blockNamenCode[0].BlockCode)                                                                             // checked ! clean and expected output
	fmt.Printf("Block Served:%s  Total:%s\n", blockSeveredAmount[0].Served, blockSeveredAmount[0].Amount)                                                                        // checked ! clean and expected output
	fmt.Printf("Cluster Name:%s  Code:%s\n", clusterNamenCode[0].ClusterName, clusterNamenCode[0].ClusterCode)                                                                   // checked ! clean and expected output
	fmt.Printf("Cluster Served:%s  Total:%s\n", clusterSeveredAmount[0].Served, clusterSeveredAmount[0].Amount)                                                                  // checked ! clean and expected output
	fmt.Printf("School UDISE : %s Name : %s Meal served : %s Amount : %s\n", schoolInfo[0].UDISECode, schoolInfo[0].SchoolName, schoolInfo[0].MealsServed, schoolInfo[0].Amount) // checked ! clean and expected output
	fmt.Printf("Ids of schools in Cluster:%s are:%v\n ", clusterNamenCode[0].ClusterName, schoolByClusteridx[0])                                                                 // checked ! clean and expected output
	fmt.Printf("The cluster:%s have schools:%v\n", schoolsByclusterInfo[0].ClusterName, schoolsByclusterInfo[0].Schools)
	fmt.Printf("Cluster totals served: %v  amount: %v\n", clusterSeveredAmount, clusterWiseAmountTotal)
	fmt.Printf("Block totals served: %v  amount: %v\n", blockWiseServedTotal, blockWiseAmountTotal)

	// ---------------------------------
	// database op
	// ---------------------------------

	db, err := database.Open("./database.db")
	if err != nil {
		fmt.Printf(":( failed to open database :%s", err)
		return
	}
	defer db.Close()

	// trying inserting

	start := time.Now()

	if err := database.InsertBlocks(db, blockNamenCode, clusterNamenCode, schoolsByclusterInfo); err != nil {
		fmt.Printf(":( failed transaction :%s", err)
	}

	fmt.Println(time.Since(start))

	// querying rows

	rows, err := db.Query(`SELECT * FROM blocks;`)
	if err != nil {
		fmt.Printf(":( failed to get rows for '(`SELECT * FROM blocks;)`' query :%s ", err)
		return
	}
	defer rows.Close()

}
