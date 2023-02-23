package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type DeleteTranscriptionVocabularyPathParams struct {
	TranscriptionVocabularyID string `pathParam:"style=simple,explode=false,name=TRANSCRIPTION_VOCABULARY_ID"`
}

type DeleteTranscriptionVocabularyRequest struct {
	PathParams DeleteTranscriptionVocabularyPathParams
	Retries    *utils.RetryConfig
}

type DeleteTranscriptionVocabularyResponse struct {
	ContentType string
	StatusCode  int
}
