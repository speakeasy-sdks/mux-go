package shared

type FilterValue struct {
	TotalCount *int64  `json:"total_count,omitempty"`
	Value      *string `json:"value,omitempty"`
}
