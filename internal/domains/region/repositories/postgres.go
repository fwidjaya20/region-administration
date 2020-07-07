package repositories

import (
	"context"
	"fmt"
	"github.com/fwidjaya20/regional-administration/internal/databases/models/province"
	"github.com/fwidjaya20/regional-administration/internal/globals"
	"github.com/jmoiron/sqlx"
	"strings"
)

type postgres struct {}

func (p *postgres) GetProvincesByCode(ctx context.Context, code []string) ([]*province.Model, error) {
	var result []*province.Model
	var err error
	var rows *sqlx.Rows

	query, args := p.buildGetProvinceByCodeQuery(code)
	rows, err = globals.GetQuery(ctx).NamedQueryxContext(ctx, query, args)

	if nil != err {
		return nil, err
	}

	for rows.Next() {
		var model province.Model
		err = rows.StructScan(&model)
		if nil != err {
			return nil, err
		}
		result = append(result, &model)
	}

	_ = rows.Close()

	return result, err
}

func (p *postgres) GetProvincesByName(ctx context.Context, name string) ([]*province.Model, error) {
	var result []*province.Model
	var err error
	var rows *sqlx.Rows

	query, args := p.buildGetProvinceByNameQuery(name)
	rows, err = globals.GetQuery(ctx).NamedQueryxContext(ctx, query, args)

	if nil != err {
		return nil, err
	}

	for rows.Next() {
		var model province.Model
		err = rows.StructScan(&model)
		if nil != err {
			return nil, err
		}
		result = append(result, &model)
	}

	_ = rows.Close()

	return result, err
}

func NewProvinceRepository() Interface {
	return &postgres{}
}

func (p *postgres) buildGetProvinceByCodeQuery(code []string) (string, interface{}) {
	var query strings.Builder
	var args = make(map[string]interface{})

	query.WriteString(fmt.Sprintf(`%s`, p.buildRegionJoinQuery()))

	if len(code) > 0 {
		query.WriteString(`WHERE p."code" IN `)
		query.WriteString(`('` + strings.Join(code, `','`) + `')`)
	}

	return query.String(), args
}

func (p *postgres) buildGetProvinceByNameQuery(name string) (string, interface{}) {
	var query strings.Builder
	var args = make(map[string]interface{})

	query.WriteString(fmt.Sprintf(`%s`, p.buildRegionJoinQuery()))
	query.WriteString(`WHERE p."name" LIKE '%` + name + `%'`)

	return query.String(), args
}

func (p *postgres) buildRegionJoinQuery() string {
	var query strings.Builder

	var columns = []string{
		`p."code" as "province_code"`,
		`p."name" as "province_name"`,
		`r."code" as "regency_code"`,
		`r."name" as "regency_name"`,
		`d."code" as "district_code"`,
		`d."name" as "district_name"`,
		`v."code" as "village_code"`,
		`v."name" as "village_name"`,
		`v."postal" as "postal_code"`,
	}

	query.WriteString(`SELECT `)
	query.WriteString(fmt.Sprintf(`%s `, strings.Join(columns, ",")))
	query.WriteString(`FROM `)
	query.WriteString(`public."provinces" p LEFT JOIN public."regencies" r ON p."code" = SUBSTRING(r."code", 1, 2) `)
	query.WriteString(`LEFT JOIN public."districts" d ON r."code" = SUBSTRING(d."code", 1, 5) `)
	query.WriteString(`LEFT JOIN public."villages" v ON d."code" = SUBSTRING(v."code", 1, 8) `)

	return query.String()
}