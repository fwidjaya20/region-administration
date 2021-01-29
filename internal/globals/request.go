package globals

type RegionalRequest struct {
	Scope         string `json:"scope"`
	WithHierarchy bool   `json:"with_hierarchy"`
}
