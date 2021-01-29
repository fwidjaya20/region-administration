package region

import (
	"context"
	"fmt"
	"github.com/fwidjaya20/regional-administration/internal/databases/models/regional"
	"github.com/fwidjaya20/regional-administration/internal/globals"
)

type ProvinceScopeStrategy struct {}

func (p ProvinceScopeStrategy) Parse(ctx context.Context, models []*regional.Model, withHierarchy bool) []globals.RegionalResponse {
	fmt.Println("[ProvinceScopeStrategy]")

	var regionalMaps = make(map[string]*globals.ProvinceResponse, 0)

	for _, model := range models {
		{
			if model.ProvinceModel.Code.Valid {
				_, isProvinceExist := regionalMaps[model.ProvinceModel.Code.String]

				if !isProvinceExist {
					regionalMaps[model.ProvinceModel.Code.String] = &globals.ProvinceResponse{
						Code: model.ProvinceModel.Code.String,
						Name: model.ProvinceModel.Name.String,
					}
				}
			}
		}

		if !withHierarchy {
			continue
		}

		// Map Regencies
		currentRegencyIndex := 0
		{
			if model.RegencyModel.Code.Valid {
				isRegencyExist := false

				for _, regency := range regionalMaps[model.ProvinceModel.Code.String].Regencies {
					if regency.Code == model.RegencyModel.Code.String {
						isRegencyExist = true
						break
					}
				}

				if !isRegencyExist {
					regionalMaps[model.ProvinceModel.Code.String].Regencies = append(regionalMaps[model.ProvinceModel.Code.String].Regencies, globals.RegencyResponse{
						Code: model.RegencyModel.Code.String,
						Name: model.RegencyModel.Name.String,
					})
				}

				currentRegencyIndex = len(regionalMaps[model.ProvinceModel.Code.String].Regencies) - 1
			}
		}

		// Map Districts
		currentDistrictIndex := 0
		{
			if model.DistrictModel.Code.Valid {
				isDistrictExist := false

				for _, district := range regionalMaps[model.ProvinceModel.Code.String].Regencies[currentRegencyIndex].Districts {
					if district.Code == model.DistrictModel.Code.String {
						isDistrictExist = true
						break
					}
				}

				if !isDistrictExist {
					regionalMaps[model.ProvinceModel.Code.String].Regencies[currentRegencyIndex].Districts = append(regionalMaps[model.ProvinceModel.Code.String].Regencies[currentRegencyIndex].Districts, globals.DistrictResponse{
						Code: model.DistrictModel.Code.String,
						Name: model.DistrictModel.Name.String,
					})
				}

				currentDistrictIndex = len(regionalMaps[model.ProvinceModel.Code.String].Regencies[currentRegencyIndex].Districts) - 1
			}
		}

		//Map Village
		{
			if model.VillageModel.Code.Valid {
				regionalMaps[model.ProvinceModel.Code.String].Regencies[currentRegencyIndex].Districts[currentDistrictIndex].Villages = append(regionalMaps[model.ProvinceModel.Code.String].Regencies[currentRegencyIndex].Districts[currentDistrictIndex].Villages, globals.VillageResponse{
					Code:       model.VillageModel.Code.String,
					Name:       model.VillageModel.Name.String,
					PostalCode: model.VillageModel.PostalCode.String,
				})
			}
		}
	}

	var result = make([]globals.RegionalResponse, 0)

	for _, v := range regionalMaps {
		result = append(result, v)
	}

	return result
}
