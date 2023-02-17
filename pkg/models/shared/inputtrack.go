package shared

type InputTrack struct {
	Channels   *int64   `json:"channels,omitempty"`
	Duration   *float64 `json:"duration,omitempty"`
	Encoding   *string  `json:"encoding,omitempty"`
	FrameRate  *float64 `json:"frame_rate,omitempty"`
	Height     *int64   `json:"height,omitempty"`
	SampleRate *int64   `json:"sample_rate,omitempty"`
	SampleSize *int64   `json:"sample_size,omitempty"`
	Type       *string  `json:"type,omitempty"`
	Width      *int64   `json:"width,omitempty"`
}
