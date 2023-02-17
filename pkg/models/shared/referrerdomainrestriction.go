package shared

// ReferrerDomainRestriction
// A list of domains allowed to play your videos.
type ReferrerDomainRestriction struct {
	AllowNoReferrer *bool    `json:"allow_no_referrer,omitempty"`
	AllowedDomains  []string `json:"allowed_domains,omitempty"`
}
