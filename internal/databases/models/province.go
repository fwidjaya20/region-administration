package models

import "database/sql"

type ProvinceModel struct {
	Code sql.NullString `db:"province_code" json:"code"`
	Name sql.NullString `db:"province_name" json:"name"`
}
