package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/utils/jwts"
)

// LogoutView 用户退出登录
// @Tags 用户管理
// @Summary 用户退出登录
// @Description 用户退出登录
// @Router /api/logout [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (*UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	token := c.Request.Header.Get("token")

	// 计算距离现在还有多久的过期时间
	err := service.ServiceApp.UserService.Logout(claims, token)

	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage("注销失败", c)
		return
	}
	res.OKWithMessage("注销成功", c)
}
