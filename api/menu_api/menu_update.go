package menu_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// MenuUpdateView 修改菜单
// @Tags 菜单管理
// @Summary 修改菜单
// @Description 修改菜单
// @Param file body models.MenuRequest true "菜单参数"
// @Router /api/menus [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (*MenuApi) MenuUpdateView(c *gin.Context) {
	var cr models.MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.LOG.Error(err)
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	id := c.Param("id")
	tx := global.DB.Begin()
	// 把之前的banner清空
	var menuModel models.MenuModel
	err = tx.Preload("Banners").Take(&menuModel, id).Error
	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage("菜单不存在", c)
		return
	}
	err = tx.Model(&menuModel).Association("Banners").Clear()
	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage("菜单图片清空失败", c)
		return
	}
	// 如果选择了banner,重新添加
	if len(cr.ImageSortList) > 0 {
		// 操作第三张表
		var bannerList []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err := tx.Create(&bannerList).Error
		if err != nil {
			res.FailWithMessage("修改菜单图片失败", c)
			return
		}
	}
	// 普通更新
	maps := structs.Map(&cr)
	err = tx.Model(&menuModel).Updates(maps).Error
	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage("修改菜单失败", c)
		return
	}
	tx.Commit()
	res.OKWithMessage("修改菜单成功", c)
}
