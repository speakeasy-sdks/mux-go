package shared

type PlaybackRestriction struct {
	CreatedAt *string                    `json:"created_at,omitempty"`
	ID        *string                    `json:"id,omitempty"`
	Referrer  *ReferrerDomainRestriction `json:"referrer,omitempty"`
	UpdatedAt *string                    `json:"updated_at,omitempty"`
}
