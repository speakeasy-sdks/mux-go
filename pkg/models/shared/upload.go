package shared

// UploadError
// Only set if an error occurred during asset creation.
type UploadError struct {
	Message *string `json:"message,omitempty"`
	Type    *string `json:"type,omitempty"`
}

type UploadStatusEnum string

const (
	UploadStatusEnumWaiting      UploadStatusEnum = "waiting"
	UploadStatusEnumAssetCreated UploadStatusEnum = "asset_created"
	UploadStatusEnumErrored      UploadStatusEnum = "errored"
	UploadStatusEnumCancelled    UploadStatusEnum = "cancelled"
	UploadStatusEnumTimedOut     UploadStatusEnum = "timed_out"
)

type Upload struct {
	AssetID          *string           `json:"asset_id,omitempty"`
	CorsOrigin       *string           `json:"cors_origin,omitempty"`
	Error            *UploadError      `json:"error,omitempty"`
	ID               *string           `json:"id,omitempty"`
	NewAssetSettings *Asset            `json:"new_asset_settings,omitempty"`
	Status           *UploadStatusEnum `json:"status,omitempty"`
	Test             *bool             `json:"test,omitempty"`
	Timeout          *int              `json:"timeout,omitempty"`
	URL              *string           `json:"url,omitempty"`
}
