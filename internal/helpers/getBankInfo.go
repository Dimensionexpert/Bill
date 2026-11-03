package helpers

import (
	"fmt"
	"strings"

	"github.com/Dimensionexpert/Bills/internal/models/data"
)

func GetBankInfo(bank [][]string) (bankinfo []data.BankInfo, err error) {

	for _, school := range bank[1:] {
		if len(school) != 15 {
			return nil, fmt.Errorf(":( unexpected number of columns: got %d, want 15", len(school))
		}
		bankinfo = append(bankinfo, data.BankInfo{
			// I made this file, so I know and I am sure bout the vlues by column,
			// so fk u n your general functions,  that works on ever bit of info that exist
			UDISE:    strings.TrimSpace(school[0]),
			Name:     strings.TrimSpace(school[1]),
			PanNo:    strings.TrimSpace(school[2]),
			TanNo:    strings.TrimSpace(school[3]),
			Address:  strings.TrimSpace(school[4]),
			City:     strings.TrimSpace(school[5]),
			Country:  strings.TrimSpace(school[6]),
			State:    strings.TrimSpace(school[7]),
			District: strings.TrimSpace(school[8]),
			PinCode:  strings.TrimSpace(school[9]),
			Mobile:   strings.TrimSpace(school[10]),
			Email:    strings.TrimSpace(school[11]),
			BankName: strings.TrimSpace(school[12]),
			IFSC:     strings.TrimSpace(school[13]),
			AccNum:   strings.TrimSpace(school[14]),
		})

	}

	return bankinfo, nil
}
