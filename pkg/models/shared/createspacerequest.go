package shared

type CreateSpaceRequest struct {
	Broadcasts  []CreateBroadcastRequest `json:"broadcasts,omitempty"`
	Passthrough *string                  `json:"passthrough,omitempty"`
	Type        *SpaceTypeEnum           `json:"type,omitempty"`
}
