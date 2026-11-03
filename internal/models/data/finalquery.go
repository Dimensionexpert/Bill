package data

import "database/sql"

type Fdata struct {
	// Block
	BlockCode string
	BlockName string

	// Cluster
	ClusterCode string
	ClusterName string

	// School
	SchoolCode string
	SchoolName sql.NullString

	// Bank information
	PanNo         sql.NullString
	TanNo         sql.NullString
	Address       sql.NullString
	City          sql.NullString
	Country       sql.NullString
	State         sql.NullString
	District      sql.NullString
	PinCode       sql.NullString
	Mobile        sql.NullString
	Email         sql.NullString
	BankName      sql.NullString
	IFSCCode      sql.NullString
	AccountNumber sql.NullString

	// Individual amounts
	CentralPrimaryAmount      sql.NullFloat64
	StatePrimaryAmount        sql.NullFloat64
	CentralUpperPrimaryAmount sql.NullFloat64
	StateUpperPrimaryAmount   sql.NullFloat64

	// Calculated totals
	PrimaryTotal      float64
	UpperPrimaryTotal float64
	GrossTotal        float64

	// Allocations
	Gen float64
	SC  float64
	ST  float64
}
