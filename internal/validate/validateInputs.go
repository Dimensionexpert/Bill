package validate

import (
	"fmt"

	"github.com/Dimensionexpert/Bills/internal/models/data"
)

func ValidateInputs(
	blockNamenCode []data.Block,
	clusterNamenCode []data.Cluster,
	schoolsByclusterInfo []data.SchoolsByCluster,
	blockServedAmount, clusterServedAmount []data.ServedAmount,
	year, month int,
) error {

	// Validate required input slices.
	if len(blockNamenCode) == 0 {
		return fmt.Errorf("block information is empty")
	}

	if len(clusterNamenCode) == 0 {
		return fmt.Errorf("cluster information is empty")
	}

	if len(blockServedAmount) == 0 {
		return fmt.Errorf("block served amounts are empty")
	}

	if len(clusterServedAmount) == 0 {
		return fmt.Errorf("cluster served amounts are empty")
	}

	if len(schoolsByclusterInfo) == 0 {
		return fmt.Errorf("school information is empty")
	}

	// Ensure block information and totals have matching lengths.
	if len(blockNamenCode) != len(blockServedAmount) {
		return fmt.Errorf(
			"block data length mismatch: headers=%d, totals=%d",
			len(blockNamenCode),
			len(blockServedAmount),
		)
	}

	// Ensure cluster information and totals have matching lengths.
	if len(clusterNamenCode) != len(clusterServedAmount) {
		return fmt.Errorf(
			"cluster data length mismatch: headers=%d, totals=%d",
			len(clusterNamenCode),
			len(clusterServedAmount),
		)
	}

	// Validate reporting period.
	if month < 1 || month > 12 {
		return fmt.Errorf("invalid month: %d", month)
	}

	if year < 2000 || year > 2050 {
		return fmt.Errorf("invalid year: %d", year)
	}

	return nil
}
