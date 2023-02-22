package operations

import (
	"github.com/speakeasy-sdks/mux-go/pkg/models/shared"
	"github.com/speakeasy-sdks/mux-go/pkg/utils"
)

type ListTranscriptionVocabulariesQueryParams struct {
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
	Page  *int `queryParam:"style=form,explode=true,name=page"`
}

type ListTranscriptionVocabulariesRequest struct {
	QueryParams ListTranscriptionVocabulariesQueryParams
	Retries     *utils.RetryConfig
}

type ListTranscriptionVocabulariesResponse struct {
	ContentType                           string
	ListTranscriptionVocabulariesResponse *shared.ListTranscriptionVocabulariesResponse
	StatusCode                            int
}
