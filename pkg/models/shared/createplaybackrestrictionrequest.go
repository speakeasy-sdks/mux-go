package shared

type CreatePlaybackRestrictionRequest struct {
	Referrer *ReferrerDomainRestriction `json:"referrer,omitempty"`
}
