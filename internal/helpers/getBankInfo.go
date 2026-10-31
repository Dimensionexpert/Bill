package helpers

import (
	"fmt"

	"github.com/Dimensionexpert/Bills/internal/models/data"
)

func GetBankInfo(bank [][]string) (bankinfo []data.Bank, err error) {

	for _, school := range bank[1:] {
		if len(school) > 0 && len(school) < 4 {
			return nil, fmt.Errorf(":( number of colums less than expected-%d : %w", len(school), err)
		}
		bankinfo = append(bankinfo, data.Bank{
			// I made this file, so I know and I am sure bout the vlues by column,
			// so fk u n your general function. that works on ever bit of info that exist
			UDISE:    school[0],
			Name:     school[1],
			AcNum:    school[2],
			UniqCode: school[3],
		})

	}

	return bankinfo, nil
}
