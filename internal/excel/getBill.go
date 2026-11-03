package excel

import (
	"fmt"

	"github.com/Dimensionexpert/Bills/internal/models/data"
	"github.com/xuri/excelize/v2"
)

func GetBill(f *excelize.File, data []data.Fdata) error {
	s := "sheet1"

	for i, school := range data {
		row := i + 2

		if err := f.SetCellValue(s, fmt.Sprintf("A%d", row), school.BlockCode); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("B%d", row), school.BlockName); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("C%d", row), school.ClusterCode); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("D%d", row), school.ClusterName); err != nil {
			return err
		}
		if err := f.SetCellValue(s, fmt.Sprintf("E%d", row), school.SchoolCode); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("F%d", row), school.SchoolName.String); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("I%d", row), school.PanNo.String); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("K%d", row), school.TanNo.String); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("N%d", row), school.Address.String); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("Q%d", row), school.City.String); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("R%d", row), school.Country.String); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("S%d", row), school.State.String); err != nil {
			return err
		}
		if err := f.SetCellValue(s, fmt.Sprintf("T%d", row), school.District.String); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("U%d", row), school.PinCode.String); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("V%d", row), school.Mobile.String); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("X%d", row), school.Email.String); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("Y%d", row), school.BankName.String); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("Z%d", row), school.IFSCCode.String); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("AA%d", row), school.AccountNumber.String); err != nil {
			return err
		}
		if err := f.SetCellValue(s, fmt.Sprintf("AB%d", row), school.CentralPrimaryAmount.Float64); err != nil {
			return err
		}
		if err := f.SetCellValue(s, fmt.Sprintf("AC%d", row), school.StatePrimaryAmount.Float64); err != nil {
			return err
		}
		if err := f.SetCellValue(s, fmt.Sprintf("AD%d", row), school.PrimaryTotal); err != nil {
			return err
		}
		if err := f.SetCellValue(s, fmt.Sprintf("AE%d", row), school.CentralUpperPrimaryAmount.Float64); err != nil {
			return err
		}
		if err := f.SetCellValue(s, fmt.Sprintf("AF%d", row), school.StateUpperPrimaryAmount.Float64); err != nil {
			return err
		}
		if err := f.SetCellValue(s, fmt.Sprintf("AG%d", row), school.UpperPrimaryTotal); err != nil {
			return err
		}
		if err := f.SetCellValue(s, fmt.Sprintf("AH%d", row), school.GrossTotal); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("AK%d", row), school.Gen); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("AL%d", row), school.SC); err != nil {
			return err
		}

		if err := f.SetCellValue(s, fmt.Sprintf("AM%d", row), school.ST); err != nil {
			return err
		}

	}

	return f.SaveAs("./Apr.xlsx")

}
