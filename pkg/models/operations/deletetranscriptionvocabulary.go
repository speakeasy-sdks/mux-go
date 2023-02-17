package operations

type DeleteTranscriptionVocabularyPathParams struct {
	TranscriptionVocabularyID string `pathParam:"style=simple,explode=false,name=TRANSCRIPTION_VOCABULARY_ID"`
}

type DeleteTranscriptionVocabularyRequest struct {
	PathParams DeleteTranscriptionVocabularyPathParams
}

type DeleteTranscriptionVocabularyResponse struct {
	ContentType string
	StatusCode  int
}
