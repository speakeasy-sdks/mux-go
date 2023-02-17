package shared

type LiveStreamLatencyModeEnum string

const (
	LiveStreamLatencyModeEnumLow      LiveStreamLatencyModeEnum = "low"
	LiveStreamLatencyModeEnumReduced  LiveStreamLatencyModeEnum = "reduced"
	LiveStreamLatencyModeEnumStandard LiveStreamLatencyModeEnum = "standard"
)

type LiveStream struct {
	ActiveAssetID              *string                               `json:"active_asset_id,omitempty"`
	AudioOnly                  *bool                                 `json:"audio_only,omitempty"`
	CreatedAt                  *string                               `json:"created_at,omitempty"`
	EmbeddedSubtitles          []LiveStreamEmbeddedSubtitleSettings  `json:"embedded_subtitles,omitempty"`
	GeneratedSubtitles         []LiveStreamGeneratedSubtitleSettings `json:"generated_subtitles,omitempty"`
	ID                         *string                               `json:"id,omitempty"`
	LatencyMode                *LiveStreamLatencyModeEnum            `json:"latency_mode,omitempty"`
	LowLatency                 *bool                                 `json:"low_latency,omitempty"`
	MaxContinuousDuration      *int                                  `json:"max_continuous_duration,omitempty"`
	NewAssetSettings           *CreateAssetRequest                   `json:"new_asset_settings,omitempty"`
	Passthrough                *string                               `json:"passthrough,omitempty"`
	PlaybackIds                []PlaybackID                          `json:"playback_ids,omitempty"`
	RecentAssetIds             []string                              `json:"recent_asset_ids,omitempty"`
	ReconnectSlateURL          *string                               `json:"reconnect_slate_url,omitempty"`
	ReconnectWindow            *float32                              `json:"reconnect_window,omitempty"`
	ReducedLatency             *bool                                 `json:"reduced_latency,omitempty"`
	SimulcastTargets           []SimulcastTarget                     `json:"simulcast_targets,omitempty"`
	Status                     *LiveStreamStatusEnum                 `json:"status,omitempty"`
	StreamKey                  *string                               `json:"stream_key,omitempty"`
	Test                       *bool                                 `json:"test,omitempty"`
	UseSlateForStandardLatency *bool                                 `json:"use_slate_for_standard_latency,omitempty"`
}
