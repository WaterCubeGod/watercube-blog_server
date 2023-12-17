package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

// SettingsInfoView 显示某一项的配置信息
func (*SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "size":
		res.OKWithData(global.CONFIG.SiteInfo, c)
	case "email":
		res.OKWithData(global.CONFIG.Email, c)
	case "qq":
		res.OKWithData(global.CONFIG.QQ, c)
	case "qi_niu":
		res.OKWithData(global.CONFIG.QiNiu, c)
	case "jwt":
		res.OKWithData(global.CONFIG.Jwt, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}

}
