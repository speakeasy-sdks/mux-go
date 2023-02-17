package shared

type Space struct {
	ActiveSessionID *string         `json:"active_session_id,omitempty"`
	Broadcasts      []Broadcast     `json:"broadcasts,omitempty"`
	CreatedAt       string          `json:"created_at"`
	ID              string          `json:"id"`
	Passthrough     *string         `json:"passthrough,omitempty"`
	Status          SpaceStatusEnum `json:"status"`
	Type            SpaceTypeEnum   `json:"type"`
}
