package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

// MenuListView 菜单列表
// @Tags 菜单管理
// @Summary 菜单列表
// @Description 菜单列表查询
// @Router /api/menus [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.MenuModel]}
func (*MenuApi) MenuListView(c *gin.Context) {

	var menuList []models.MenuModel
	var menuIDList []uint
	global.DB.Order("sort asc").Find(&menuList).Select("id").Scan(&menuIDList)

	// 查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort asc").Find(&menuBanners, "menu_id in ?", menuIDList)

	var menus []MenuResponse
	for _, menu := range menuList {
		var banners []Banner
		for _, banner := range menuBanners {
			if menu.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: menu,
			Banners:   banners,
		})
	}
	res.OKWithData(menus, c)
}
