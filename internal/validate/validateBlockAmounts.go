package validate

import (
	"fmt"
	"strconv"

	"github.com/Dimensionexpert/Bills/internal/models/data"
)

func GetTotalAmountsBlock(
	clusterSeveredAmount []data.ServedAmount,
) (blockWiseServedTotal []int, blockWiseAmountTotal []int, err error) {
	// since block is only 1 element here ( each pdf only have 1 block) bt i am so so genius i took edge case in hand  and thought of some hyporthetical situation where pdf can have more than 1 block and hence  stored it in []. Also, fortunately they don't have [][]struct thing the loop is easy
	// i have to sum the clusterwise served and amount total and check it against the csv sourced totals for block.

	var blockWiseTotalServed []int
	var blockWiseTotalAmount []int

	servedTotal := 0
	amountTotal := 0

	for i, cluster := range clusterSeveredAmount {
		served, err := strconv.Atoi(cluster.Served)
		if err != nil {
			return nil, nil, fmt.Errorf(
				"invalid served value %q at cluster index %d: %w",
				cluster.Served, i, err,
			)
		}

		amount, err := strconv.Atoi(cluster.Amount)
		if err != nil {
			return nil, nil, fmt.Errorf(
				"invalid amount value %q at cluster index %d: %w",
				cluster.Amount, i, err,
			)
		}

		servedTotal += served
		amountTotal += amount
	}

	blockWiseTotalServed = append(blockWiseTotalServed, servedTotal)
	blockWiseTotalAmount = append(blockWiseTotalAmount, amountTotal)

	return blockWiseTotalServed, blockWiseTotalAmount, nil
}

func ValidateBlockTotals(
	blockSeveredAmount []data.ServedAmount,
	blockWiseServedTotal, blockWiseAmountTotal []int,
) (bool, error) {
	// using same approch as ValidateClusterTotals() since the structure is same

	if len(blockSeveredAmount) != len(blockWiseServedTotal) ||
		len(blockSeveredAmount) != len(blockWiseAmountTotal) {
		return false, fmt.Errorf(
			"block totals length mismatch: source=%d, calculated served=%d, calculated amount=%d",
			len(blockSeveredAmount),
			len(blockWiseServedTotal),
			len(blockWiseAmountTotal),
		)
	}

	for i := range blockSeveredAmount {
		served, err := strconv.Atoi(blockSeveredAmount[i].Served)
		if err != nil {
			return false, fmt.Errorf(
				"invalid served value %q at block index %d: %w",
				blockSeveredAmount[i].Served, i, err,
			)
		}

		amount, err := strconv.Atoi(blockSeveredAmount[i].Amount)
		if err != nil {
			return false, fmt.Errorf(
				"invalid amount value %q at block index %d: %w",
				blockSeveredAmount[i].Amount, i, err,
			)
		}

		if blockWiseServedTotal[i] != served {
			return false, fmt.Errorf(
				"served mismatch at block index %d: calculated %d, source %d",
				i, blockWiseServedTotal[i], served,
			)
		}

		if blockWiseAmountTotal[i] != amount {
			return false, fmt.Errorf(
				"amount mismatch at block index %d: calculated %d, source %d",
				i, blockWiseAmountTotal[i], amount,
			)
		}
	}

	return true, nil
}
