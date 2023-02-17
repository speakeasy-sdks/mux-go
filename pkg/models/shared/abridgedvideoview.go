package shared

type AbridgedVideoView struct {
	CountryCode           *string `json:"country_code,omitempty"`
	ErrorTypeID           *int    `json:"error_type_id,omitempty"`
	ID                    *string `json:"id,omitempty"`
	PlayerErrorCode       *string `json:"player_error_code,omitempty"`
	PlayerErrorMessage    *string `json:"player_error_message,omitempty"`
	TotalRowCount         *int64  `json:"total_row_count,omitempty"`
	VideoTitle            *string `json:"video_title,omitempty"`
	ViewEnd               *string `json:"view_end,omitempty"`
	ViewStart             *string `json:"view_start,omitempty"`
	ViewerApplicationName *string `json:"viewer_application_name,omitempty"`
	ViewerOsFamily        *string `json:"viewer_os_family,omitempty"`
}
