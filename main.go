package main

import (
	"fmt"
	"time"

	"github.com/Dimensionexpert/Bills/internal/database"

	"github.com/Dimensionexpert/Bills/internal/excel"
	"github.com/Dimensionexpert/Bills/internal/helpers"
	"github.com/xuri/excelize/v2"
)

func main() {
	// ---------------------------------
	// open the csv
	// ---------------------------------
	record, err := helpers.Open("./source/Sep/c68.csv")
	if err != nil {
		fmt.Printf("can't fking open the file %s", err)
		return
	}

	bank, err := helpers.Open("./source/sources/bank.csv")
	if err != nil {
		fmt.Printf("can't fking open the file %s", err)
		return
	}

	template, err := excelize.OpenFile("templates/template.xlsx")
	if err != nil {
		fmt.Printf("can't open the template %s", err)
	}
	defer template.Close()

	rows, err := template.GetRows("Sheet1")
	if err != nil {
		fmt.Printf("can't get the rows template %s", err)
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

	// // ---------------------------------
	// // cluster wise totals
	// // ---------------------------------
	// clusterWiseServedTotal, clusterWiseAmountTotal, err := validate.GetTotalAmountsCluster(schoolsByclusterInfo)
	// if err != nil {
	// 	fmt.Printf(":( calculating cluster totals failed: %v\n", err)
	// }

	// clusterOk, err := validate.ValidateClusterTotals(
	// 	clusterSeveredAmount,
	// 	clusterWiseServedTotal,
	// 	clusterWiseAmountTotal,
	// )

	// if err != nil {
	// 	fmt.Printf(":( Cluster validation failed: %v\n", err)
	// 	return
	// }

	// if !clusterOk {
	// 	fmt.Println(":( Cluster validation failed")
	// 	return
	// }

	// fmt.Println(":) Cluster validation passed")

	// // ---------------------------------
	// // block wise validation
	// // ---------------------------------

	// blockWiseServedTotal, blockWiseAmountTotal, err := validate.GetTotalAmountsBlock(clusterSeveredAmount)
	// if err != nil {
	// 	fmt.Printf(":( calculating block totals failed: %v\n", err)
	// }

	// blockOk, err := validate.ValidateBlockTotals(blockSeveredAmount, blockWiseServedTotal, blockWiseAmountTotal)
	// if err != nil {
	// 	fmt.Printf(":( Block validation failed: %v\n", err)
	// 	return
	// }

	// if !blockOk {
	// 	fmt.Println(":( Block validation failed")
	// 	return
	// }

	// fmt.Println(":) Block validation passed")

	// BANK INFO
	bankInfo, err := helpers.GetBankInfo(bank)
	if err != nil {
		fmt.Printf(":( error getting bank info : %s", err)
		return
	}

	// ---------------------------------
	// data cross checking
	// ---------------------------------

	fmt.Printf("Block Name:%s  Code:%s\n", blockNamenCode[0].BlockName, blockNamenCode[0].BlockCode)                                                                             // checked ! clean and expected output
	fmt.Printf("Block Served:%s  Total:%s\n", blockSeveredAmount[0].Served, blockSeveredAmount[0].Amount)                                                                        // checked ! clean and expected output
	fmt.Printf("Cluster Name:%s  Code:%s\n", clusterNamenCode[0].ClusterName, clusterNamenCode[0].ClusterCode)                                                                   // checked ! clean and expected output
	fmt.Printf("Cluster Served:%s  Total:%s\n", clusterSeveredAmount[0].Served, clusterSeveredAmount[0].Amount)                                                                  // checked ! clean and expected output
	fmt.Printf("School UDISE : %s Name : %s Meal served : %s Amount : %s\n", schoolInfo[0].UDISECode, schoolInfo[0].SchoolName, schoolInfo[0].MealsServed, schoolInfo[0].Amount) // checked ! clean and expected output
	fmt.Printf("The clusterCode:%s clusterName:%s have schools:%v\n", schoolsByclusterInfo[0].CLusterCode, schoolsByclusterInfo[0].ClusterName, schoolsByclusterInfo[0].Schools)
	// fmt.Printf("Cluster totals served: %v  amount: %v\n", clusterSeveredAmount, clusterWiseAmountTotal)
	// fmt.Printf("Block totals served: %v  amount: %v\n", blockWiseServedTotal, blockWiseAmountTotal)

	// ---------------------------------
	// database op
	// ---------------------------------

	// this variables (year month section scope) are temporarily declared here, i want this as input from user.
	// altho i am thinking since the the section and sope parsing is pretty consistent from the csv i can leave it like this,
	// so user have less inputs.
	year := 2026
	month := 9
	section, scope, err := helpers.GetSectionScope(record)

	if err != nil {
		fmt.Printf(":( cant resolve section and scope cause: %s\n", err)
	}

	db, err := database.Open("./database.db")
	if err != nil {
		fmt.Printf(":( failed to open database :%s", err)
		return
	}
	defer db.Close()

	// trying inserting

	start := time.Now()

	if err := database.InsertIntoDB(db, blockNamenCode, bankInfo, clusterNamenCode, schoolsByclusterInfo, blockSeveredAmount,
		clusterSeveredAmount, scope, section, year, month); err != nil {
		fmt.Printf(":( failed transaction :%s", err)
	}

	if err := database.InsertBankInfo(db, bankInfo); err != nil {
		fmt.Printf(":( failed tranaction %s\n", err)
	}

	fmt.Println(time.Since(start))

	data, err := database.GetFinal(db, year, month)
	if err != nil {
		fmt.Printf("DUCK !! %s", err)
	}

	fmt.Println(data[0])
	fmt.Println(rows[0])
	excel.GetBill(template, data)

}
