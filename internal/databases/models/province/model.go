package province

import "github.com/fwidjaya20/regional-administration/internal/databases/models"

type Model struct {
	models.ProvinceModel
	models.RegencyModel
	models.DistrictModel
	models.VillageModel
}
