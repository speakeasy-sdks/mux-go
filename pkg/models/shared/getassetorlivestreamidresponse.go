package shared

type GetAssetOrLiveStreamIDResponseDataObjectTypeEnum string

const (
	GetAssetOrLiveStreamIDResponseDataObjectTypeEnumAsset      GetAssetOrLiveStreamIDResponseDataObjectTypeEnum = "asset"
	GetAssetOrLiveStreamIDResponseDataObjectTypeEnumLiveStream GetAssetOrLiveStreamIDResponseDataObjectTypeEnum = "live_stream"
)

// GetAssetOrLiveStreamIDResponseDataObject
// Describes the Asset or LiveStream object associated with the playback ID.
type GetAssetOrLiveStreamIDResponseDataObject struct {
	ID   *string                                           `json:"id,omitempty"`
	Type *GetAssetOrLiveStreamIDResponseDataObjectTypeEnum `json:"type,omitempty"`
}

type GetAssetOrLiveStreamIDResponseData struct {
	ID     *string                                   `json:"id,omitempty"`
	Object *GetAssetOrLiveStreamIDResponseDataObject `json:"object,omitempty"`
	Policy *PlaybackPolicyEnum                       `json:"policy,omitempty"`
}

type GetAssetOrLiveStreamIDResponse struct {
	Data *GetAssetOrLiveStreamIDResponseData `json:"data,omitempty"`
}
