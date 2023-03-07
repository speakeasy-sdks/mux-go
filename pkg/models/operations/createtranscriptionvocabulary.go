package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type CreateTranscriptionVocabularyRequest struct {
	Request shared.CreateTranscriptionVocabularyRequest `request:"mediaType=application/json"`
}

type CreateTranscriptionVocabularyResponse struct {
	ContentType                     string
	StatusCode                      int
	RawResponse                     *http.Response
	TranscriptionVocabularyResponse *shared.TranscriptionVocabularyResponse
}
