package region

import (
	"github.com/fwidjaya20/regional-administration/internal/databases/models/regional"
	"github.com/fwidjaya20/regional-administration/internal/globals"
)

func MapToProvince(memo map[string]*globals.ProvinceResponse, model *regional.Model) {
	if model.ProvinceModel.Code.Valid {
		_, isProvinceExist := memo[model.ProvinceModel.Code.String]

		if !isProvinceExist {
			memo[model.ProvinceModel.Code.String] = &globals.ProvinceResponse{
				Code: model.ProvinceModel.Code.String,
				Name: model.ProvinceModel.Name.String,
			}
		}
	}
}

func MapToRegency(memo []globals.RegencyResponse, model *regional.Model) (int, []globals.RegencyResponse) {
	if model.RegencyModel.Code.Valid {
		isRegencyExist := false

		for _, regency := range memo {
			if regency.Code == model.RegencyModel.Code.String {
				isRegencyExist = true
				break
			}
		}

		if !isRegencyExist {
			memo = append(memo, globals.RegencyResponse{
				Code: model.RegencyModel.Code.String,
				Name: model.RegencyModel.Name.String,
			})
		}

		return len(memo) - 1, memo
	}

	return 0, []globals.RegencyResponse{}
}

func MapToDistrict(memo []globals.DistrictResponse, model *regional.Model) (int, []globals.DistrictResponse) {
	if model.DistrictModel.Code.Valid {
		isDistrictExist := false

		for _, district := range memo {
			if district.Code == model.DistrictModel.Code.String {
				isDistrictExist = true
				break
			}
		}

		if !isDistrictExist {
			memo = append(memo, globals.DistrictResponse{
				Code: model.DistrictModel.Code.String,
				Name: model.DistrictModel.Name.String,
			})
		}

		return len(memo) - 1, memo
	}

	return 0, []globals.DistrictResponse{}
}