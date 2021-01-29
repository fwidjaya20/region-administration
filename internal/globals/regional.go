package globals

type RegionalResponse interface {}

type ProvinceResponse struct {
	RegionalResponse `json:"-"`
	Code             string            `json:"code"`
	Name             string            `json:"name"`
	Regencies        []RegencyResponse `json:"regencies,omitempty"`
}

type RegencyResponse struct {
	RegionalResponse `json:"-"`
	Code             string             `json:"code"`
	Name             string             `json:"name"`
	Province         *ProvinceResponse  `json:"province,omitempty"`
	Districts        []DistrictResponse `json:"districts,omitempty"`
}

type DistrictResponse struct {
	RegionalResponse `json:"-"`
	Code             string            `json:"code"`
	Name             string            `json:"name"`
	Regency          *RegencyResponse  `json:"regency,omitempty"`
	Villages         []VillageResponse `json:"villages,omitempty"`
}

type VillageResponse struct {
	RegionalResponse `json:"-"`
	Code             string            `json:"code"`
	Name             string            `json:"name"`
	PostalCode       string            `json:"postal_code"`
	District         *DistrictResponse `json:"district,omitempty"`
}
