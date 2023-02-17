package shared

type Error struct {
	Code        *int64   `json:"code,omitempty"`
	Count       *int64   `json:"count,omitempty"`
	Description *string  `json:"description,omitempty"`
	ID          *int64   `json:"id,omitempty"`
	LastSeen    *string  `json:"last_seen,omitempty"`
	Message     *string  `json:"message,omitempty"`
	Notes       *string  `json:"notes,omitempty"`
	Percentage  *float64 `json:"percentage,omitempty"`
}
