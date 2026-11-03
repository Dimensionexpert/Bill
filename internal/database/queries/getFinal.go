package queries

import (
	"database/sql"
	"fmt"

	_ "embed"

	"github.com/Dimensionexpert/Bills/internal/models/data"
)

//go:embed getAll.sql
var query string

func GetFinal(db *sql.DB, year, month int) ([]data.Fdata, error) {

	rows, err := db.Query(
		query,
		year, month, // STATE PRIMARY
		year, month, // STATE UPPER_PRIMARY
		year, month, // CENTRAL PRIMARY
		year, month, // CENTRAL UPPER_PRIMARY
	)
	if err != nil {
		return nil, fmt.Errorf(":( error scanning rows %s", err)
	}
	var results []data.Fdata

	for rows.Next() {
		var d data.Fdata
		err := rows.Scan(
			&d.BlockCode,
			&d.BlockName,
			&d.ClusterCode,
			&d.ClusterName,
			&d.SchoolCode,
			&d.SchoolName,

			&d.PanNo,
			&d.TanNo,
			&d.Address,
			&d.City,
			&d.Country,
			&d.State,
			&d.District,
			&d.PinCode,
			&d.Mobile,
			&d.Email,
			&d.BankName,
			&d.IFSCCode,
			&d.AccountNumber,

			&d.CentralPrimaryAmount,
			&d.StatePrimaryAmount,
			&d.PrimaryTotal,

			&d.CentralUpperPrimaryAmount,
			&d.StateUpperPrimaryAmount,
			&d.UpperPrimaryTotal,

			&d.GrossTotal,
			&d.Gen,
			&d.SC,
			&d.ST,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan report row: %w", err)
		}
		defer rows.Close()

		results = append(results, d)

	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate report rows: %w", err)
	}

	return results, nil
}
