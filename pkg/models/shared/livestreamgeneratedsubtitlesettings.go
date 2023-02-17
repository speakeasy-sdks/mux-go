package shared

type LiveStreamGeneratedSubtitleSettingsLanguageCodeEnum string

const (
	LiveStreamGeneratedSubtitleSettingsLanguageCodeEnumEn   LiveStreamGeneratedSubtitleSettingsLanguageCodeEnum = "en"
	LiveStreamGeneratedSubtitleSettingsLanguageCodeEnumEnUs LiveStreamGeneratedSubtitleSettingsLanguageCodeEnum = "en-US"
)

type LiveStreamGeneratedSubtitleSettings struct {
	LanguageCode               *LiveStreamGeneratedSubtitleSettingsLanguageCodeEnum `json:"language_code,omitempty"`
	Name                       *string                                              `json:"name,omitempty"`
	Passthrough                *string                                              `json:"passthrough,omitempty"`
	TranscriptionVocabularyIds []string                                             `json:"transcription_vocabulary_ids,omitempty"`
}
