package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// MenuDetailView 菜单细节
// @Tags 菜单管理
// @Summary 菜单细节
// @Description 菜单细节查询
// @Param id query int true "查询菜单的id"
// @Router /api/menu_detail [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (*MenuApi) MenuDetailView(c *gin.Context) {
	id := c.Param("id")
	var menuModel models.MenuModel
	err := global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	// 查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", id)

	banners := make([]Banner, 0)
	for _, banner := range menuBanners {
		if menuModel.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.BannerID,
			Path: banner.BannerModel.Path,
		})
	}
	menu := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	res.OKWithData(menu, c)
}
