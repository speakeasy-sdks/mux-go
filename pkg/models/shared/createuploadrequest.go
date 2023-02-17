package shared

type CreateUploadRequest struct {
	CorsOrigin       *string            `json:"cors_origin,omitempty"`
	NewAssetSettings CreateAssetRequest `json:"new_asset_settings"`
	Test             *bool              `json:"test,omitempty"`
	Timeout          *int               `json:"timeout,omitempty"`
}
