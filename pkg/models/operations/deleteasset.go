package operations

type DeleteAssetPathParams struct {
	AssetID string `pathParam:"style=simple,explode=false,name=ASSET_ID"`
}

type DeleteAssetRequest struct {
	PathParams DeleteAssetPathParams
}

type DeleteAssetResponse struct {
	ContentType string
	StatusCode  int
}
