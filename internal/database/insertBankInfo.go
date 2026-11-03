package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Dimensionexpert/Bills/internal/helpers"
	"github.com/Dimensionexpert/Bills/internal/models/data"
)

func InsertBankInfo(db *sql.DB, bankInfo []data.BankInfo) error {
	// starting context for timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// starting transaction to maintain atomicity
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(":( failed to begin transaction: %w", err)
	}
	defer tx.Rollback()
	insertBankInfo := `
    INSERT INTO school_banks (
        school_code,
        school_name,
        pan_num,
        tan_num,
        address,
        city,
        country,
        state,
        district,
        pincode,
        mobile,
        email,
        bank_name,
        ifsc_code,
        account_number
    )
    VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
    ON CONFLICT (school_code)
    DO UPDATE SET
        school_name = excluded.school_name,
        pan_num = excluded.pan_num,
        tan_num = excluded.tan_num,
        address = excluded.address,
        city = excluded.city,
        country = excluded.country,
        state = excluded.state,
        district = excluded.district,
        pincode = excluded.pincode,
        mobile = excluded.mobile,
        email = excluded.email,
        bank_name = excluded.bank_name,
        ifsc_code = excluded.ifsc_code,
        account_number = excluded.account_number
	`
	insertBankInfoStmt, err := tx.PrepareContext(ctx, insertBankInfo)
	if err != nil {
		return fmt.Errorf(":( failed to prepare insertBankInfo statement: %w", err)
	}
	defer insertBankInfoStmt.Close()

	for _, school := range bankInfo {
		if _, err := insertBankInfoStmt.ExecContext(
			ctx,
			school.UDISE,
			school.Name,
			helpers.Nullable(school.PanNo),
			helpers.Nullable(school.TanNo),
			helpers.Nullable(school.Address),
			helpers.Nullable(school.City),
			helpers.Nullable(school.Country),
			helpers.Nullable(school.State),
			helpers.Nullable(school.District),
			helpers.Nullable(school.PinCode),
			helpers.Nullable(school.Mobile),
			helpers.Nullable(school.Email),
			helpers.Nullable(school.BankName),
			helpers.Nullable(school.IFSC),
			helpers.Nullable(school.AccNum),
		); err != nil {
			return fmt.Errorf(":( failed to insert bank Info for school code -%s school-%s pan :%s tan :%s :%w", school.UDISE, school.Name, school.PanNo, school.TanNo, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf(":( failed to commit bank info transaction: %w", err)
	}

	return nil

}
