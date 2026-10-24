package helpers

import (
	"fmt"
	"strings"

	"github.com/Dimensionexpert/Bills/internal/models/data"
)

func GetServedAmount(record [][]string, index []int) ([]data.Data, error) {
	var info []data.Data

	for _, idx := range index {

		_, served, sesep := strings.Cut(record[idx][2], "::")
		if !sesep {
			return nil, fmt.Errorf(
				"name: no separator found in %q at row %d",
				record[idx][2], idx,
			)
		}
		_, amount, amsep := strings.Cut(record[idx][4], "::")
		if !amsep {
			return nil, fmt.Errorf(
				"code: no separator found in %q at row %d",
				record[idx][0], idx,
			)
		}

		info = append(info, data.Data{
			First:  strings.TrimSpace(served), // total served comes first in csv
			Second: strings.TrimSpace(amount), // total amount is the second
		})
	}

	return info, nil
}
