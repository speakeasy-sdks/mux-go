package shared

type DeliveryReportAssetStateEnum string

const (
	DeliveryReportAssetStateEnumReady   DeliveryReportAssetStateEnum = "ready"
	DeliveryReportAssetStateEnumErrored DeliveryReportAssetStateEnum = "errored"
	DeliveryReportAssetStateEnumDeleted DeliveryReportAssetStateEnum = "deleted"
)

// DeliveryReportDeliveredSecondsByResolution
// Seconds delivered broken into resolution tiers. Each tier will only be displayed if there was content delivered in the tier.
type DeliveryReportDeliveredSecondsByResolution struct {
	Tier1080p     *float64 `json:"tier_1080p,omitempty"`
	Tier720p      *float64 `json:"tier_720p,omitempty"`
	TierAudioOnly *float64 `json:"tier_audio_only,omitempty"`
}

type DeliveryReport struct {
	AssetDuration                *float64                                    `json:"asset_duration,omitempty"`
	AssetID                      *string                                     `json:"asset_id,omitempty"`
	AssetState                   *DeliveryReportAssetStateEnum               `json:"asset_state,omitempty"`
	CreatedAt                    *string                                     `json:"created_at,omitempty"`
	DeletedAt                    *string                                     `json:"deleted_at,omitempty"`
	DeliveredSeconds             *float64                                    `json:"delivered_seconds,omitempty"`
	DeliveredSecondsByResolution *DeliveryReportDeliveredSecondsByResolution `json:"delivered_seconds_by_resolution,omitempty"`
	LiveStreamID                 *string                                     `json:"live_stream_id,omitempty"`
	Passthrough                  *string                                     `json:"passthrough,omitempty"`
}
