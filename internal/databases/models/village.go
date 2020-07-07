package models

import "database/sql"

type VillageModel struct {
	Code       sql.NullString `db:"village_code" json:"code"`
	Name       sql.NullString `db:"village_name" json:"name"`
	PostalCode sql.NullString `db:"postal_code" json:"postal_code"`
}
