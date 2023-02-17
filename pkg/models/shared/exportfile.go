package shared

type ExportFile struct {
	Path    *string `json:"path,omitempty"`
	Type    *string `json:"type,omitempty"`
	Version *int    `json:"version,omitempty"`
}
