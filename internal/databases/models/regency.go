package models

import "database/sql"

type RegencyModel struct {
	Code sql.NullString `db:"regency_code" json:"code"`
	Name sql.NullString `db:"regency_name" json:"name"`
}
