package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
)

type GetTranscriptionVocabularyPathParams struct {
	TranscriptionVocabularyID string `pathParam:"style=simple,explode=false,name=TRANSCRIPTION_VOCABULARY_ID"`
}

type GetTranscriptionVocabularyRequest struct {
	PathParams GetTranscriptionVocabularyPathParams
}

type GetTranscriptionVocabularyResponse struct {
	ContentType                     string
	StatusCode                      int
	TranscriptionVocabularyResponse *shared.TranscriptionVocabularyResponse
}
