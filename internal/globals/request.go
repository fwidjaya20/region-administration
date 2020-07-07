package globals

type RegionalRequest struct {
	Scope string   `json:"scope"`
	Code  []string `json:"code"`
	Name  string   `json:"name"`
}
