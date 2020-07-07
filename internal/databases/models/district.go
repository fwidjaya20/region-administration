package models

import "database/sql"

type DistrictModel struct {
	Code sql.NullString `db:"district_code" json:"code"`
	Name sql.NullString `db:"district_name" json:"name"`
}
