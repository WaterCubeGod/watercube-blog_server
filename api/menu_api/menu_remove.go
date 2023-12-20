package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// MenuRemoveView 删除菜单
// @Tags 菜单管理
// @Summary 删除菜单
// @Description 删除菜单
// @Param menus body models.MenuId true "删除菜单的列表"
// @Router /api/menus [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (*MenuApi) MenuRemoveView(c *gin.Context) {
	var menuId models.MenuId
	err := c.ShouldBindJSON(&menuId)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var menuList []models.MenuModel
	tx := global.DB.Begin()
	count := global.DB.Find(&menuList, menuId.ID).RowsAffected
	for _, menuModel := range menuList {

		// 删除关联图片
		err = tx.Model(&menuModel).Association("Banners").Clear()
		if err != nil {
			global.LOG.Error(err)
			res.FailWithMessage(fmt.Sprintf("删除与%v相关联图片失败", menuModel.Title), c)
			return
		}

		// 删除菜单
		err = tx.Delete(&menuModel).Error
		if err != nil {
			global.LOG.Error(err)
			res.FailWithMessage(fmt.Sprintf("删除%v菜单失败", menuModel.Title), c)
			return
		}
	}
	tx.Commit()
	res.OKWithMessage(fmt.Sprintf("共删除%d个菜单", count), c)
}
