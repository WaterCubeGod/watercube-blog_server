package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

// MenuNameListView 菜单名称列表
// @Tags 菜单管理
// @Summary 菜单名称列表
// @Description 菜单名称列表
// @Router /api/menu_names [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.MenuModel]}
func (*MenuApi) MenuNameListView(c *gin.Context) {
	var menuName []MenuNameResponse
	global.DB.Model(models.MenuModel{}).
		Select("id", "id", "title", "path").Scan(&menuName)
	res.OKWithData(menuName, c)
}
