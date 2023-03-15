package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type CreateTranscriptionVocabularyResponse struct {
	ContentType                     string
	StatusCode                      int
	RawResponse                     *http.Response
	TranscriptionVocabularyResponse *shared.TranscriptionVocabularyResponse
}
