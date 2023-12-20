package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwts"
	"gvb_server/utils/pwd"
)

type PwdRequest struct {
	OldPwd string `json:"old_pwd"`
	Pwd    string `json:"pwd"`
}

// UserUpdatePwd 用户修改信息
// @Tags 用户管理
// @Summary 用户修改信息
// @Description 用户修改信息
// @Param data body string false "用户密码"
// @Router /api/user_rule [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (*UserApi) UserUpdatePwd(c *gin.Context) {
	var cr PwdRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	if !pwd.ComparePasswords(user.Password, cr.OldPwd) {
		res.FailWithMessage("旧密码不正确", c)
		return
	}
	newPwd := pwd.HashEncode(cr.Pwd)
	err = global.DB.Model(&user).Update("password", newPwd).Error
	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage("密码更新失败", c)
		return
	}
	res.OKWithMessage("密码更新成功", c)
}
