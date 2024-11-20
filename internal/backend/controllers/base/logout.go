package base

import (
	"net/http"

	"github.com/An-Yan-d/ChineseSubFinder/pkg/types/backend"

	"github.com/An-Yan-d/ChineseSubFinder/pkg/common"
	"github.com/gin-gonic/gin"
)

func (cb *ControllerBase) LogoutHandler(c *gin.Context) {

	// 注销
	common.SetAccessToken("")
	c.JSON(http.StatusOK, backend.ReplyCommon{Message: "ok, need ReLogin"})
}
