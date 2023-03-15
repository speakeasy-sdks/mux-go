package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type GetTranscriptionVocabularyRequest struct {
	TranscriptionVocabularyID string `pathParam:"style=simple,explode=false,name=TRANSCRIPTION_VOCABULARY_ID"`
}

type GetTranscriptionVocabularyResponse struct {
	ContentType                     string
	StatusCode                      int
	RawResponse                     *http.Response
	TranscriptionVocabularyResponse *shared.TranscriptionVocabularyResponse
}
