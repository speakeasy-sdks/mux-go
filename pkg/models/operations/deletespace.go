package operations

type DeleteSpacePathParams struct {
	SpaceID string `pathParam:"style=simple,explode=false,name=SPACE_ID"`
}

type DeleteSpaceRequest struct {
	PathParams DeleteSpacePathParams
}

type DeleteSpaceResponse struct {
	ContentType string
	StatusCode  int
}
