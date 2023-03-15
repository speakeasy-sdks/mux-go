package operations

import (
	"net/http"
)

type DeleteTranscriptionVocabularyRequest struct {
	TranscriptionVocabularyID string `pathParam:"style=simple,explode=false,name=TRANSCRIPTION_VOCABULARY_ID"`
}

type DeleteTranscriptionVocabularyResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
