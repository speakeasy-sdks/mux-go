package shared

type TranscriptionVocabulary struct {
	CreatedAt   *string  `json:"created_at,omitempty"`
	ID          *string  `json:"id,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Passthrough *string  `json:"passthrough,omitempty"`
	Phrases     []string `json:"phrases,omitempty"`
	UpdatedAt   *string  `json:"updated_at,omitempty"`
}
