package shared

type InputFile struct {
	ContainerFormat *string      `json:"container_format,omitempty"`
	Tracks          []InputTrack `json:"tracks,omitempty"`
}
