package validate

import (
	"fmt"
	"strconv"

	"github.com/Dimensionexpert/Bills/internal/models/data"
)

func GetTotalAmountsCluster(schoolsByclusterInfo []data.SchoolsByCluster) (clusterWiseServedTotal []int, clusterWiseAmountTotal []int, err error) {

	var clusterWiseTotalAmount []int
	var clusterWiseTotalServed []int
	for i := range schoolsByclusterInfo {
		AmountlTotal := 0
		ServedTotal := 0

		school := schoolsByclusterInfo[i].Schools

		for j := range school {
			amount, err := strconv.Atoi(school[j].Amount)
			if err != nil {
				return nil, nil, fmt.Errorf(
					"invalid amount for school %s in cluster %s: %w",
					school[j].UDISECode,
					schoolsByclusterInfo[i].ClusterName,
					err,
				)
			}

			served, err := strconv.Atoi(school[j].MealsServed)
			if err != nil {
				return nil, nil, fmt.Errorf(
					"invalid meals served for school %s in cluster %s: %w",
					school[j].UDISECode,
					schoolsByclusterInfo[i].ClusterName,
					err,
				)
			}
			AmountlTotal += amount
			ServedTotal += served
		}
		clusterWiseTotalAmount = append(clusterWiseTotalAmount, AmountlTotal)
		clusterWiseTotalServed = append(clusterWiseTotalServed, ServedTotal)

	}

	return clusterWiseTotalServed, clusterWiseTotalAmount, nil
}

func ValidateClusterTotals(clusterServedAmount []data.ServedAmount, clusterWiseServedTotal, clusterWiseAmountTotal []int) (bool, error) {
	// since clusterServedAmount stores a struct that holds Served, Amount respectively
	// {2377 3684} like this. here {Served:2377 Amount:3684} as string and this data is extracted from csv.
	// we comparing this with our totals calculated by GetTotalAmountsCluster()
	// if len(clusterServedAmount) != len(clusterWiseServedTotal) ||
	// 	len(clusterServedAmount) != len(clusterWiseAmountTotal) {
	// 	return false, fmt.Errorf(
	// 		"cluster totals length mismatch: source=%d, calculated served=%d, calculated amount=%d",
	// 		len(clusterServedAmount),
	// 		len(clusterWiseServedTotal),
	// 		len(clusterWiseAmountTotal),
	// 	)
	// }
	for i := range clusterWiseServedTotal {

		served, err := strconv.Atoi(clusterServedAmount[i].Served)
		if err != nil {
			return false, fmt.Errorf(
				"invalid served value %q at cluster idx %d: %w",
				clusterServedAmount[i].Served, i, err,
			)
		}

		amount, err := strconv.Atoi(clusterServedAmount[i].Amount)
		if err != nil {
			return false, fmt.Errorf(
				"invalid amount value %q at cluster idx %d: %w",
				clusterServedAmount[i].Amount, i, err,
			)
		}

		if clusterWiseServedTotal[i] != served {
			return false, fmt.Errorf(
				"served mismatch at cluster idx %d: calculated %d, got %d",
				i, clusterWiseServedTotal[i], served,
			)
		}

		if clusterWiseAmountTotal[i] != amount {
			return false, fmt.Errorf(
				"amount mismatch at cluster idx %d: calculated %d, got %d",
				i, clusterWiseAmountTotal[i], amount,
			)
		}
	}

	return true, nil
}
