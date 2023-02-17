package shared

type Incident struct {
	AffectedViews              *int64                     `json:"affected_views,omitempty"`
	AffectedViewsPerHour       *int64                     `json:"affected_views_per_hour,omitempty"`
	AffectedViewsPerHourOnOpen *int64                     `json:"affected_views_per_hour_on_open,omitempty"`
	Breakdowns                 []IncidentBreakdown        `json:"breakdowns,omitempty"`
	Description                *string                    `json:"description,omitempty"`
	ErrorDescription           *string                    `json:"error_description,omitempty"`
	ID                         *string                    `json:"id,omitempty"`
	Impact                     *string                    `json:"impact,omitempty"`
	IncidentKey                *string                    `json:"incident_key,omitempty"`
	MeasuredValue              *float64                   `json:"measured_value,omitempty"`
	MeasuredValueOnClose       *float64                   `json:"measured_value_on_close,omitempty"`
	Measurement                *string                    `json:"measurement,omitempty"`
	NotificationRules          []IncidentNotificationRule `json:"notification_rules,omitempty"`
	Notifications              []IncidentNotification     `json:"notifications,omitempty"`
	ResolvedAt                 *string                    `json:"resolved_at,omitempty"`
	SampleSize                 *int64                     `json:"sample_size,omitempty"`
	SampleSizeUnit             *string                    `json:"sample_size_unit,omitempty"`
	Severity                   *string                    `json:"severity,omitempty"`
	StartedAt                  *string                    `json:"started_at,omitempty"`
	Status                     *string                    `json:"status,omitempty"`
	Threshold                  *float64                   `json:"threshold,omitempty"`
}
