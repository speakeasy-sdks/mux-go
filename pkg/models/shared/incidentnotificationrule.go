package shared

type IncidentNotificationRule struct {
	Action     *string            `json:"action,omitempty"`
	ID         *string            `json:"id,omitempty"`
	PropertyID *string            `json:"property_id,omitempty"`
	Rules      []NotificationRule `json:"rules,omitempty"`
	Status     *string            `json:"status,omitempty"`
}
