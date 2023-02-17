package shared

type LiveStreamEmbeddedSubtitleSettingsLanguageChannelEnum string

const (
	LiveStreamEmbeddedSubtitleSettingsLanguageChannelEnumCc1 LiveStreamEmbeddedSubtitleSettingsLanguageChannelEnum = "cc1"
)

type LiveStreamEmbeddedSubtitleSettings struct {
	LanguageChannel *LiveStreamEmbeddedSubtitleSettingsLanguageChannelEnum `json:"language_channel,omitempty"`
	LanguageCode    *string                                                `json:"language_code,omitempty"`
	Name            *string                                                `json:"name,omitempty"`
	Passthrough     *string                                                `json:"passthrough,omitempty"`
}
