package shared

type UpdateAssetMasterAccessRequestMasterAccessEnum string

const (
	UpdateAssetMasterAccessRequestMasterAccessEnumTemporary UpdateAssetMasterAccessRequestMasterAccessEnum = "temporary"
	UpdateAssetMasterAccessRequestMasterAccessEnumNone      UpdateAssetMasterAccessRequestMasterAccessEnum = "none"
)

type UpdateAssetMasterAccessRequest struct {
	MasterAccess *UpdateAssetMasterAccessRequestMasterAccessEnum `json:"master_access,omitempty"`
}
