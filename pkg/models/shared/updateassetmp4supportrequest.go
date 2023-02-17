package shared

type UpdateAssetMp4SupportRequestMp4SupportEnum string

const (
	UpdateAssetMp4SupportRequestMp4SupportEnumStandard UpdateAssetMp4SupportRequestMp4SupportEnum = "standard"
	UpdateAssetMp4SupportRequestMp4SupportEnumNone     UpdateAssetMp4SupportRequestMp4SupportEnum = "none"
)

type UpdateAssetMp4SupportRequest struct {
	Mp4Support *UpdateAssetMp4SupportRequestMp4SupportEnum `json:"mp4_support,omitempty"`
}
