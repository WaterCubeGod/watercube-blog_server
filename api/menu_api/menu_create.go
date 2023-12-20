package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// MenuCreateView 添加菜单
// @Tags 菜单管理
// @Summary 添加菜单
// @Description 添加菜单
// @Param file body models.MenuRequest true "添加菜单"
// @Router /api/menus [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (*MenuApi) MenuCreateView(c *gin.Context) {
	tx := global.DB.Begin()
	var cr models.MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, "msg", c)
		return
	}
	var menuList []models.MenuModel
	count := tx.Find(&menuList, "title = ? or path = ?", cr.Title, cr.Path).RowsAffected
	if count > 0 {
		res.FailWithMessage("重复的菜单", c)
		return
	}

	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}

	// 创建banner数据入库
	err = tx.Create(&menuModel).Error
	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage("菜单添加失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.OKWithMessage("菜单添加成功", c)
		return
	}
	var menuBannerList []models.MenuBannerModel
	for _, sort := range cr.ImageSortList {
		// 查看是否真的有这张图片
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	// 给第三张表入库
	err = tx.Create(&menuBannerList).Error
	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage("菜单图片关联失败", c)
		return
	}
	tx.Commit()
	res.OKWithMessage("添加成功", c)
}
