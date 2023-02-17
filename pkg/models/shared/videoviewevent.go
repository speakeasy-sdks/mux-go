package shared

type VideoViewEvent struct {
	EventTime    *int64  `json:"event_time,omitempty"`
	Name         *string `json:"name,omitempty"`
	PlaybackTime *int64  `json:"playback_time,omitempty"`
	ViewerTime   *int64  `json:"viewer_time,omitempty"`
}
