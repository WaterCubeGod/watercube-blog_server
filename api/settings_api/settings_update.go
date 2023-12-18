package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

// SettingsUpdateView 修改配置信息
// @Tags 配置管理
// @Summary 修改配置信息
// @Description 修改某一项配置信息
// @Param data body SettingsUri false "表示单个参数"
// @Router /api/settings/:name [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (s *SettingsApi) SettingsUpdateView(c *gin.Context) {
	cr := SettingsUri{}
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "size":
		var info config.SiteInfo
		err = c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.CONFIG.SiteInfo = info
	case "email":
		var email config.Email
		err = c.ShouldBindJSON(&email)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.CONFIG.Email = email
	case "qq":
		var qq config.QQ
		err = c.ShouldBindJSON(&qq)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.CONFIG.QQ = qq
	case "qi_niu":
		var qiNiu config.Telegraph
		err = c.ShouldBindJSON(&qiNiu)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.CONFIG.QiNiu = qiNiu
	case "jwt":
		var jwt config.Jwt
		err = c.ShouldBindJSON(&jwt)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.CONFIG.Jwt = jwt
	default:
		res.FailWithMessage("没有对应的配置信息", c)
		return
	}
	s.SettingsUpdate(c)
}

func (*SettingsApi) SettingsUpdate(c *gin.Context) {
	err := core.SetYaml()
	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage(err.Error(), c)
	}
	res.OKWithMessage("系统信息修改成功", c)
}
