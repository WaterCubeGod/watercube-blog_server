package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OKWithData(global.CONFIG.SiteInfo, c)
}
