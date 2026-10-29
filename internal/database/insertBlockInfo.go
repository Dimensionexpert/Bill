package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Dimensionexpert/Bills/internal/models/data"
)

func InsertBlocks(db *sql.DB, blockNamenCode []data.Block, clusterNamenCode []data.Cluster, schoolsByclusterInfo []data.SchoolsByCluster) error {

	// create context here
	// 
	// benchmarking the insertion
	// re-import (skips on conflict): 1.222123ms
	// fresh insert: 7.147733ms
	// 5 seconds timeout is more than enough for now, as this parser only parses one file at a time.
	// currently only block, cluster and school codes and names are being inserted.
	// the meals served and amount data from each school's report is not yet being inserted.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// start the transaction here
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf(":( failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// inserting the Blocks
	InsertBlock := "INSERT INTO blocks (block_code,block_name) VALUES (?,?) ON CONFLICT DO NOTHING;"
	blockStmt, err := tx.PrepareContext(ctx, InsertBlock)
	if err != nil {
		return fmt.Errorf(":( failed to prepare InsertBlock statement: %w", err)
	}
	defer blockStmt.Close()

	for _, block := range blockNamenCode {
		if _, err := blockStmt.ExecContext(ctx, block.BlockCode, block.BlockName); err != nil {
			return fmt.Errorf(":( failed to insert blockCode %s, blockName %s: %w", block.BlockCode, block.BlockName, err)
		}
	}

	// inserting the cluster
	InsertCluster := "INSERT INTO clusters (block_code,cluster_code,cluster_name) VALUES (?,?,?) ON CONFLICT DO NOTHING;"
	clusterStmt, err := tx.PrepareContext(ctx, InsertCluster)
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

	// inserting schools
	InsertSchool := "INSERT INTO schools (cluster_code,school_code,school_name) VALUES (?,?,?) ON CONFLICT DO NOTHING;"
	schoolStmt, err := tx.PrepareContext(ctx, InsertSchool)
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

	// commiting the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf(":( failed to commit transaction: %w", err)
	}

	return nil

}
