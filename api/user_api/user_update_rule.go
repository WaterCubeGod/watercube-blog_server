package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
)

type UserRole struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	NickName string     `json:"nick_name"`
	UserID   uint       `json:"user_id" binding:"required" msg:"用户id错误"`
}

// UserUpdateRoleView 管理员修改用户信息
// @Tags 用户管理
// @Summary 管理员修改用户信息
// @Description 管理员修改用户信息
// @Param data body UserRole false "用户信息"
// @Router /api/user_rule [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (*UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserRole
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, "msg", c)
		return
	}
	var user models.UserModel
	err = global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.FailWithMessage("不存在该用户", c)
		return
	}
	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error
	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage("用户信息更新失败", c)
		return
	}
	res.FailWithMessage("用户信息更新成功", c)
}
