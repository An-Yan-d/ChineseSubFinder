package base

import (
	"net/http"

	"github.com/An-Yan-d/ChineseSubFinder/pkg"

	"github.com/An-Yan-d/ChineseSubFinder/pkg/settings"
	"github.com/gin-gonic/gin"
)

func (cb *ControllerBase) DefSettingsHandler(c *gin.Context) {
	var err error
	defer func() {
		// 统一的异常处理
		cb.ErrorProcess(c, "DefSettingsHandler", err)
	}()

	c.JSON(http.StatusOK, settings.NewSettings(pkg.ConfigRootDirFPath()))
}
