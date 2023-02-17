package shared

type UpdateTranscriptionVocabularyRequest struct {
	Name        *string  `json:"name,omitempty"`
	Passthrough *string  `json:"passthrough,omitempty"`
	Phrases     []string `json:"phrases"`
}
