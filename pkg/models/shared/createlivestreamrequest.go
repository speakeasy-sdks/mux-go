package shared

type CreateLiveStreamRequestLatencyModeEnum string

const (
	CreateLiveStreamRequestLatencyModeEnumLow      CreateLiveStreamRequestLatencyModeEnum = "low"
	CreateLiveStreamRequestLatencyModeEnumReduced  CreateLiveStreamRequestLatencyModeEnum = "reduced"
	CreateLiveStreamRequestLatencyModeEnumStandard CreateLiveStreamRequestLatencyModeEnum = "standard"
)

type CreateLiveStreamRequest struct {
	AudioOnly                  *bool                                   `json:"audio_only,omitempty"`
	EmbeddedSubtitles          []LiveStreamEmbeddedSubtitleSettings    `json:"embedded_subtitles,omitempty"`
	GeneratedSubtitles         []LiveStreamGeneratedSubtitleSettings   `json:"generated_subtitles,omitempty"`
	LatencyMode                *CreateLiveStreamRequestLatencyModeEnum `json:"latency_mode,omitempty"`
	LowLatency                 *bool                                   `json:"low_latency,omitempty"`
	MaxContinuousDuration      *int                                    `json:"max_continuous_duration,omitempty"`
	NewAssetSettings           *CreateAssetRequest                     `json:"new_asset_settings,omitempty"`
	Passthrough                *string                                 `json:"passthrough,omitempty"`
	PlaybackPolicy             []PlaybackPolicyEnum                    `json:"playback_policy,omitempty"`
	ReconnectSlateURL          *string                                 `json:"reconnect_slate_url,omitempty"`
	ReconnectWindow            *float32                                `json:"reconnect_window,omitempty"`
	ReducedLatency             *bool                                   `json:"reduced_latency,omitempty"`
	SimulcastTargets           []CreateSimulcastTargetRequest          `json:"simulcast_targets,omitempty"`
	Test                       *bool                                   `json:"test,omitempty"`
	UseSlateForStandardLatency *bool                                   `json:"use_slate_for_standard_latency,omitempty"`
}
