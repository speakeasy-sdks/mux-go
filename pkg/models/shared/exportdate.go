package shared

import (
	"github.com/speakeasy-sdks/mux-go/pkg/types"
)

type ExportDate struct {
	ExportDate *types.Date  `json:"export_date,omitempty"`
	Files      []ExportFile `json:"files,omitempty"`
}
