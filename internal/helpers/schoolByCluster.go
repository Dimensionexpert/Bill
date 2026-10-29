package helpers

import (
	"fmt"

	"github.com/Dimensionexpert/Bills/internal/models/data"
)

func SchoolsByClusterIdx(clusterCodeNameidx, clusterServedAmountidx, schoolIdx []int) ([][]int, error) {
	var schoolsByCluster [][]int // store slices of schools in that
	for i := range clusterCodeNameidx {
		var school []int // store the idx of that schools in that iteration
		for _, value := range schoolIdx {
			if value > clusterCodeNameidx[i] && value < clusterServedAmountidx[i] {
				school = append(school, value)
			}
		}

		schoolsByCluster = append(schoolsByCluster, school)

	}

	return schoolsByCluster, nil

}

func SchoolsByClusterinfo(record [][]string, clusterinfo []data.Cluster, index [][]int) ([]data.SchoolsByCluster, error) {
	var SchoolsByCluster []data.SchoolsByCluster
	for i := range index {
		schools, err := GetSchools(record, index[i])
		if err != nil {
			return []data.SchoolsByCluster{}, fmt.Errorf("Error getting school information")
		}

		SchoolsByCluster = append(SchoolsByCluster, data.SchoolsByCluster{
			CLusterCode: clusterinfo[i].ClusterCode,
			ClusterName: clusterinfo[i].ClusterName,
			Schools:     schools,
		})
	}

	return SchoolsByCluster, nil
}
