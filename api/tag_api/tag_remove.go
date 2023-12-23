package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// TagRemoveView 标签删除
// @Tags  标签管理
// @Summary 标签删除
// @Description 标签删除
// @Param menus body models.RemoveRequest true "删除标签的列表"
// @Router /api/tags [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (*TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 查看是否有文章引用标签
	var tags []models.TagModel
	count := global.DB.Find(&tags, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("没有需要产出的标签", c)
		return
	}
	num := 0
	tx := global.DB.Begin()
	for _, tag := range tags {
		count := tx.Find(&tag).Association("Articles").Count()
		if count != 0 {
			continue
		}
		err := tx.Delete(&tag).Error
		if err != nil {
			global.LOG.Error(err)
			res.FailWithMessage("删除标签失败", c)
			return
		}
		num++
	}
	tx.Commit()
	res.OKWithMessage(fmt.Sprintf("共删除%d个标签", num), c)
}
