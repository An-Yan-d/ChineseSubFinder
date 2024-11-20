package backend

import (
	"github.com/An-Yan-d/ChineseSubFinder/pkg/settings"
)

type ReqSetupInfo struct {
	Settings settings.Settings `json:"settings" binding:"required"`
}
