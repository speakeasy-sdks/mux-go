package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/models/utils"
)

type CreateTranscriptionVocabularyRequest struct {
	Request shared.CreateTranscriptionVocabularyRequest `request:"mediaType=application/json"`
	Retries *utils.RetryConfig
}

type CreateTranscriptionVocabularyResponse struct {
	ContentType                     string
	StatusCode                      int
	TranscriptionVocabularyResponse *shared.TranscriptionVocabularyResponse
}
