package helpers

import (
	"fmt"
	"strings"

	"github.com/Dimensionexpert/Bills/internal/models/data"
)

func GetClusterCodeNameIDs(record [][]string) ([]int, error) {
	var index []int

	for i, row := range record {
		if len(row) == 0 {
			continue
		}

		if strings.Contains(strings.ToLower(strings.TrimSpace(row[0])), "cluster code") {
			index = append(index, i)
		}
	}

	return index, nil
}

func GetClusterInfo(record [][]string, index []int) ([]data.Cluster, error) {
	var clusters []data.Cluster
	clusterInfo, err := GetCodeName(record, index)
	if err != nil {
		return nil, fmt.Errorf("error getting the code and name of cluster:%w", err)
	}
	for i, _ := range index {
		clusters = append(clusters, data.Cluster{
			ClusterCode: clusterInfo[i].First,
			ClusterName: clusterInfo[i].Second,
		})
	}

	return clusters, nil

}

func GetClusterServedAmountIDs(record [][]string) ([]int, error) {
	var index []int

	for i, row := range record {
		if len(row) == 0 {
			continue
		}

		if strings.Contains(strings.ToLower(strings.TrimSpace(row[0])), "cluster wise total") {
			index = append(index, i)
		}
	}

	return index, nil
}

func GetClusterServedAmount(record [][]string, index []int) ([]data.ServedAmount, error) {
	var clustersTotals []data.ServedAmount
	clusterInfo, err := GetServedAmount(record, index)
	if err != nil {
		return nil, fmt.Errorf("error getting the code and name of cluster:%w", err)
	}
	for i, _ := range index {
		clustersTotals = append(clustersTotals, data.ServedAmount{
			Served: clusterInfo[i].First,
			Amount: clusterInfo[i].Second,
		})
	}

	return clustersTotals, nil

}
