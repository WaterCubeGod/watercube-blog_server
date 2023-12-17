package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

func (SettingsApi) SettingsUpdateView(c *gin.Context) {
	cr := config.SiteInfo{}
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}
	global.CONFIG.SiteInfo = cr
	err = core.SetYaml()
	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage(err.Error(), c)
	}
	res.OKWithMessage("系统信息修改成功", c)
}
