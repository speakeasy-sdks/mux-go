package shared

type TrackStatusEnum string

const (
	TrackStatusEnumPreparing TrackStatusEnum = "preparing"
	TrackStatusEnumReady     TrackStatusEnum = "ready"
	TrackStatusEnumErrored   TrackStatusEnum = "errored"
	TrackStatusEnumDeleted   TrackStatusEnum = "deleted"
)

type TrackTextSourceEnum string

const (
	TrackTextSourceEnumUploaded           TrackTextSourceEnum = "uploaded"
	TrackTextSourceEnumEmbedded           TrackTextSourceEnum = "embedded"
	TrackTextSourceEnumGeneratedLive      TrackTextSourceEnum = "generated_live"
	TrackTextSourceEnumGeneratedLiveFinal TrackTextSourceEnum = "generated_live_final"
)

type TrackTextTypeEnum string

const (
	TrackTextTypeEnumSubtitles TrackTextTypeEnum = "subtitles"
)

type TrackTypeEnum string

const (
	TrackTypeEnumVideo TrackTypeEnum = "video"
	TrackTypeEnumAudio TrackTypeEnum = "audio"
	TrackTypeEnumText  TrackTypeEnum = "text"
)

type Track struct {
	ClosedCaptions   *bool                `json:"closed_captions,omitempty"`
	Duration         *float64             `json:"duration,omitempty"`
	ID               *string              `json:"id,omitempty"`
	LanguageCode     *string              `json:"language_code,omitempty"`
	MaxChannelLayout *string              `json:"max_channel_layout,omitempty"`
	MaxChannels      *int64               `json:"max_channels,omitempty"`
	MaxFrameRate     *float64             `json:"max_frame_rate,omitempty"`
	MaxHeight        *int64               `json:"max_height,omitempty"`
	MaxWidth         *int64               `json:"max_width,omitempty"`
	Name             *string              `json:"name,omitempty"`
	Passthrough      *string              `json:"passthrough,omitempty"`
	Status           *TrackStatusEnum     `json:"status,omitempty"`
	TextSource       *TrackTextSourceEnum `json:"text_source,omitempty"`
	TextType         *TrackTextTypeEnum   `json:"text_type,omitempty"`
	Type             *TrackTypeEnum       `json:"type,omitempty"`
}
