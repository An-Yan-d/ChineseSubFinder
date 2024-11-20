package backend

import "github.com/An-Yan-d/ChineseSubFinder/pkg/settings"

type ReqCheckProxy struct {
	ProxySettings settings.ProxySettings `json:"proxy_settings"  binding:"required"`
}
