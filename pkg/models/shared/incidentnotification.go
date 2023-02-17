package shared

type IncidentNotification struct {
	AttemptedAt *string `json:"attempted_at,omitempty"`
	ID          *int64  `json:"id,omitempty"`
	QueuedAt    *string `json:"queued_at,omitempty"`
}
