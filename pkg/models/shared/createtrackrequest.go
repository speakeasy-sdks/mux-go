package shared

type CreateTrackRequestTextTypeEnum string

const (
	CreateTrackRequestTextTypeEnumSubtitles CreateTrackRequestTextTypeEnum = "subtitles"
)

type CreateTrackRequestTypeEnum string

const (
	CreateTrackRequestTypeEnumText CreateTrackRequestTypeEnum = "text"
)

type CreateTrackRequest struct {
	ClosedCaptions *bool                          `json:"closed_captions,omitempty"`
	LanguageCode   string                         `json:"language_code"`
	Name           *string                        `json:"name,omitempty"`
	Passthrough    *string                        `json:"passthrough,omitempty"`
	TextType       CreateTrackRequestTextTypeEnum `json:"text_type"`
	Type           CreateTrackRequestTypeEnum     `json:"type"`
	URL            string                         `json:"url"`
}
