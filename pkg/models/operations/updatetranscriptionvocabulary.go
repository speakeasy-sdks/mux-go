package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
	"net/http"
)

type UpdateTranscriptionVocabularyPathParams struct {
	TranscriptionVocabularyID string `pathParam:"style=simple,explode=false,name=TRANSCRIPTION_VOCABULARY_ID"`
}

type UpdateTranscriptionVocabularyRequest struct {
	PathParams UpdateTranscriptionVocabularyPathParams
	Request    shared.UpdateTranscriptionVocabularyRequest `request:"mediaType=application/json"`
	Retries    *utils.RetryConfig
}

type UpdateTranscriptionVocabularyResponse struct {
	ContentType                     string
	StatusCode                      int
	RawResponse                     *http.Response
	TranscriptionVocabularyResponse *shared.TranscriptionVocabularyResponse
}
