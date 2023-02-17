package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type CreateTranscriptionVocabularyRequest struct {
	Request shared.CreateTranscriptionVocabularyRequest `request:"mediaType=application/json"`
}

type CreateTranscriptionVocabularyResponse struct {
	ContentType                     string
	StatusCode                      int
	TranscriptionVocabularyResponse *shared.TranscriptionVocabularyResponse
}
