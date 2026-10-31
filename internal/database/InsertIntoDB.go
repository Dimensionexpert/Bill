package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Dimensionexpert/Bills/internal/models/data"
	"github.com/Dimensionexpert/Bills/internal/validate"
)

func InsertIntoDB(db *sql.DB, blockNamenCode []data.Block, bankInfo []data.Bank, clusterNamenCode []data.Cluster, schoolsByclusterInfo []data.SchoolsByCluster, blockServedAmount, clusterServedAmount []data.ServedAmount, scope, section string, year, month int) error {

	// create context here
	//
	// benchmarking the insertion
	// re-import (skips on conflict): 1.384061ms
	// fresh insert: 5.792153ms
	// 5 seconds timeout is more than enough for now, as this parser only parses one file at a time.

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// validating inputs before starting transaction.
	if err := validate.ValidateInputs(blockNamenCode, clusterNamenCode, schoolsByclusterInfo, blockServedAmount, clusterServedAmount, year, month); err != nil {
		return fmt.Errorf("input validation failed: %w", err)
	}

	// start the transaction here
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(":( failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// inserting the Block info
	insertBlock := `
		INSERT INTO blocks (
			block_code,
			block_name
		)
	 	VALUES (?,?)
		ON CONFLICT DO NOTHING;`
	blockStmt, err := tx.PrepareContext(ctx, insertBlock)
	if err != nil {
		return fmt.Errorf(":( failed to prepare InsertBlock statement: %w", err)
	}
	defer blockStmt.Close()

	for _, block := range blockNamenCode {
		if _, err := blockStmt.ExecContext(ctx, block.BlockCode, block.BlockName); err != nil {
			return fmt.Errorf(":( failed to insert blockCode %s, blockName %s: %w", block.BlockCode, block.BlockName, err)
		}
	}

	// inserting the cluster info
	insertCluster := `
		INSERT INTO clusters (
			block_code,
			cluster_code,
			cluster_name
		)
		VALUES (?,?,?)
		ON CONFLICT DO NOTHING;
	`
	clusterStmt, err := tx.PrepareContext(ctx, insertCluster)
	if err != nil {
		return fmt.Errorf(":( failed to prepare InsertCluster statemente: %w", err)
	}
	defer clusterStmt.Close()

	for _, cluster := range clusterNamenCode {
		// using blockNamenCode[0].BlockCode its a temparory thing, asuming this whole parser will only parse only one month of a block.
		// if i ever thought  this parser will parse more than one block at a time i'll fix this
		if _, err := clusterStmt.ExecContext(ctx, blockNamenCode[0].BlockCode, cluster.ClusterCode, cluster.ClusterName); err != nil {
			return fmt.Errorf(":( failed to insert clusterCode %s, clusterName %s: %w", cluster.ClusterCode, cluster.ClusterName, err)
		}
	}

	// inserting schools info
	insertSchool := `
		INSERT INTO schools (
			cluster_code,
			school_code,
			school_name
		)
		VALUES (?,?,?)
		ON CONFLICT DO NOTHING;
	`
	schoolStmt, err := tx.PrepareContext(ctx, insertSchool)
	if err != nil {
		return fmt.Errorf(":( failed to prepare InsertSchool statemente: %w", err)
	}
	defer schoolStmt.Close()

	for _, cluster := range schoolsByclusterInfo {
		clusterCode := cluster.CLusterCode
		for _, school := range cluster.Schools {
			if _, err := schoolStmt.ExecContext(ctx, clusterCode, school.UDISECode, school.SchoolName); err != nil {
				return fmt.Errorf(":( failed to insert clusterCode %s, schoolCode %s, schoolName %s: %w", clusterCode, school.UDISECode, school.SchoolName, err)

			}
		}
	}

	// inserting the bank infos
	insertBank := `
		INSERT INTO banks (
			unique_code,
			school_code,
			school_name,
			acount_num
		)
		VALUES (?,?,?,?)
		ON CONFLICT DO NOTHING
	`
	// setting it as ON CONFLICT DO NOTHING until i figure out what to do with it
	insertBankStmt, err := tx.PrepareContext(ctx, insertBank)
	if err != nil {
		return fmt.Errorf(":( failed to prepare InsertBank statemente: %w", err)
	}
	defer insertBankStmt.Close()

	for _, bank := range bankInfo {
		if _, err := insertBankStmt.ExecContext(ctx, bank.UniqCode, bank.UDISE, bank.Name, bank.AcNum); err != nil {
			return fmt.Errorf(":( insertBankStmt: failed to insert schoolCode %s, schoolName %s: %w", bank.UDISE, bank.Name, err)
		}
	}

	// using the ON CONFLICT DO NOTHING for info tables, i,e blocks, clusters, schools assuming,
	// the names rarely changes if name changed i have to DML query to update names of concerned entity.
	// using the ON CONFLICT DO UPDATE for report tabls, i.e block_reports, cluster_reports, school_reports as they,
	// hold the cruicial and sensitive info that might need to be updated more frequently than entity names.

	// inserting the block totals
	insertBlockTotal := `
	    INSERT INTO block_reports (
	        scope, section, year, month, block_code,
	        total_served, total_amount
	    )
	    VALUES (?, ?, ?, ?, ?, ?, ?)
	    ON CONFLICT(scope, section, year, month, block_code)
	    DO UPDATE SET
	        total_served = excluded.total_served,
	        total_amount = excluded.total_amount;
	`
	blockTotalStmt, err := tx.PrepareContext(ctx, insertBlockTotal)
	if err != nil {
		return fmt.Errorf(":( failed to prepare Blocktotal statemente: %w", err)
	}
	defer blockTotalStmt.Close()

	for i, block := range blockServedAmount {
		// if len(blockNamenCode) == len(blockSeveredAmount){}
		// not checking boundry cause validate functions validates, correctness of data already.
		blockCode := blockNamenCode[i].BlockCode

		if _, err := blockTotalStmt.ExecContext(ctx, scope, section, year, month, blockCode, block.Served, block.Amount); err != nil {
			return fmt.Errorf(":( failed to insert blockTotal %s :%w", blockCode, err)
		}
	}

	// inserting the cluster totals
	insertClusterTotal := `
	    INSERT INTO cluster_reports (
	        scope, section, year, month, cluster_code,
	        total_served, total_amount
	    )
	    VALUES (?, ?, ?, ?, ?, ?, ?)
	    ON CONFLICT(scope, section, year, month, cluster_code)
	    DO UPDATE SET
	        total_served = excluded.total_served,
	        total_amount = excluded.total_amount;
	`
	clusterTotalStmt, err := tx.PrepareContext(ctx, insertClusterTotal)
	if err != nil {
		return fmt.Errorf(":( failed to prepare Clustertotal statemente: %w", err)
	}
	defer clusterTotalStmt.Close()

	for i, cluster := range clusterServedAmount {
		clusterCode := clusterNamenCode[i].ClusterCode
		if _, err := clusterTotalStmt.ExecContext(ctx, scope, section, year, month, clusterCode, cluster.Served, cluster.Amount); err != nil {
			return fmt.Errorf(":( failed to insert ClusterTotal %s :%w", clusterCode, err)
		}
	}

	// inserting the school totals
	insertSchoolsTotal := `
	    INSERT INTO school_reports (
	        scope, section, year, month, school_code,
	        served, amount
	    )
	    VALUES (?, ?, ?, ?, ?, ?, ?)
	    ON CONFLICT(scope, section, year, month, school_code)
	    DO UPDATE SET
	        served = excluded.served,
	        amount = excluded.amount;
	`
	insertSchoolsStmt, err := tx.PrepareContext(ctx, insertSchoolsTotal)
	if err != nil {
		return fmt.Errorf(":( failed to prepare Schoolstotal statemente: %w", err)
	}
	defer insertSchoolsStmt.Close()

	for _, cluster := range schoolsByclusterInfo {
		for _, school := range cluster.Schools {
			if _, err := insertSchoolsStmt.ExecContext(ctx, scope, section, year, month, school.UDISECode, school.MealsServed, school.Amount); err != nil {
				return fmt.Errorf(":( failed to insert schoolCode %s schoolName %s :%w", school.UDISECode, school.SchoolName, err)
			}
		}
	}

	// commiting the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf(":( failed to commit transaction: %w", err)
	}

	return nil

}
