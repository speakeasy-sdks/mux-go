package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type GetTranscriptionVocabularyPathParams struct {
	TranscriptionVocabularyID string `pathParam:"style=simple,explode=false,name=TRANSCRIPTION_VOCABULARY_ID"`
}

type GetTranscriptionVocabularyRequest struct {
	PathParams GetTranscriptionVocabularyPathParams
	Retries    *utils.RetryConfig
}

type GetTranscriptionVocabularyResponse struct {
	ContentType                     string
	StatusCode                      int
	TranscriptionVocabularyResponse *shared.TranscriptionVocabularyResponse
}
