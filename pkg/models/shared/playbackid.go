package shared

type PlaybackID struct {
	ID     *string             `json:"id,omitempty"`
	Policy *PlaybackPolicyEnum `json:"policy,omitempty"`
}
