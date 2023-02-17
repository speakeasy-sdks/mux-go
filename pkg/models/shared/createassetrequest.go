package shared

type CreateAssetRequestMasterAccessEnum string

const (
	CreateAssetRequestMasterAccessEnumNone      CreateAssetRequestMasterAccessEnum = "none"
	CreateAssetRequestMasterAccessEnumTemporary CreateAssetRequestMasterAccessEnum = "temporary"
)

type CreateAssetRequestMp4SupportEnum string

const (
	CreateAssetRequestMp4SupportEnumNone     CreateAssetRequestMp4SupportEnum = "none"
	CreateAssetRequestMp4SupportEnumStandard CreateAssetRequestMp4SupportEnum = "standard"
)

type CreateAssetRequest struct {
	Input          []InputSettings                     `json:"input,omitempty"`
	MasterAccess   *CreateAssetRequestMasterAccessEnum `json:"master_access,omitempty"`
	Mp4Support     *CreateAssetRequestMp4SupportEnum   `json:"mp4_support,omitempty"`
	NormalizeAudio *bool                               `json:"normalize_audio,omitempty"`
	Passthrough    *string                             `json:"passthrough,omitempty"`
	PerTitleEncode *bool                               `json:"per_title_encode,omitempty"`
	PlaybackPolicy []PlaybackPolicyEnum                `json:"playback_policy,omitempty"`
	Test           *bool                               `json:"test,omitempty"`
}
