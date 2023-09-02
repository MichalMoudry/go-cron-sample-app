package queries

import _ "embed"

var (
	//go:embed scripts/queries/GetJobResults.sql
	GetJobResults string
)
