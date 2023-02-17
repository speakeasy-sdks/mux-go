package shared

type UpdateLiveStreamEmbeddedSubtitlesRequest struct {
	EmbeddedSubtitles []LiveStreamEmbeddedSubtitleSettings `json:"embedded_subtitles,omitempty"`
}
