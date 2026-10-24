package helpers

import (
	"fmt"
	"strings"

	"github.com/Dimensionexpert/Bills/internal/models/data"
)

func GetCodeName(record [][]string, index []int) ([]data.Data, error) {
	var info []data.Data

	for _, idx := range index {
		_, code, cosep := strings.Cut(record[idx][0], "::")
		if !cosep {
			return nil, fmt.Errorf(
				"code: no separator found in %q at row %d",
				record[idx][0], idx,
			)
		}

		_, name, nasep := strings.Cut(record[idx][2], "::")
		if !nasep {
			return nil, fmt.Errorf(
				"name: no separator found in %q at row %d",
				record[idx][2], idx,
			)
		}

		info = append(info, data.Data{
			First:  strings.TrimSpace(code), // code comes first in csv
			Second: strings.TrimSpace(name), // name comes second
		})
	}

	return info, nil
}

