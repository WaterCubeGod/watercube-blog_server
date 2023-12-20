package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// UserRemoveView 用户注销
// @Tags 用户管理
// @Summary 用户注销
// @Description 用户注销
// @Param data body models.RemoveRequest true "用户id列表"
// @Router /api/users [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (*UserApi) UserRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var users []models.UserModel
	count := global.DB.Find(&users, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("您选择的用户不存在", c)
		return
	}
	tx := global.DB.Begin()
	// TODO:删除用户收藏，用户发布的文章

	err = tx.Delete(&users).Error
	if err != nil {
		res.FailWithMessage("删除失败", c)
		return
	}
	tx.Commit()
	res.OKWithMessage(fmt.Sprintf("共删除%d个用户", count), c)
}
