package query

import _ "embed"

var (
	//go:embed queries/CreateTables.sql
	CreateTables string
)
