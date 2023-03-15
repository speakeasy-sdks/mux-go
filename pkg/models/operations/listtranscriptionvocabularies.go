package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListTranscriptionVocabulariesRequest struct {
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
	Page  *int `queryParam:"style=form,explode=true,name=page"`
}

type ListTranscriptionVocabulariesResponse struct {
	ContentType                           string
	ListTranscriptionVocabulariesResponse *shared.ListTranscriptionVocabulariesResponse
	StatusCode                            int
	RawResponse                           *http.Response
}
