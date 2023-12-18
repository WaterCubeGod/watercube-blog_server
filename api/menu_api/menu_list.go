package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// MenuListView 菜单列表
// @Tags 菜单管理
// @Summary 菜单列表
// @Description 菜单列表查询
// @Router /api/menus [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.MenuModel]}
func (*MenuApi) MenuListView(c *gin.Context) {
	var menuModel []models.MenuModel
	global.DB.First(&menuModel)
	res.OKWithData(menuModel, c)
}
