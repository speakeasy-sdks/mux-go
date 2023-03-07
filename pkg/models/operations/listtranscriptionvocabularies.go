package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"net/http"
)

type ListTranscriptionVocabulariesQueryParams struct {
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
	Page  *int `queryParam:"style=form,explode=true,name=page"`
}

type ListTranscriptionVocabulariesRequest struct {
	QueryParams ListTranscriptionVocabulariesQueryParams
}

type ListTranscriptionVocabulariesResponse struct {
	ContentType                           string
	ListTranscriptionVocabulariesResponse *shared.ListTranscriptionVocabulariesResponse
	StatusCode                            int
	RawResponse                           *http.Response
}
