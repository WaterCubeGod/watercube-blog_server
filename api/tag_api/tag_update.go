package tag_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// TagUpdateView 标签更新
// @Tags 标签管理
// @Summary 标签更新
// @Description 标签更新
// @Param file body models.TagModel true "标签参数"
// @Router /api/tag [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (*TagApi) TagUpdateView(c *gin.Context) {
	var cr models.TagModel
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var tag models.TagModel
	err = global.DB.Take(&tag, cr.ID).Error
	if err != nil {
		res.FailWithMessage("标签不存在", c)
		return
	}
	err = global.DB.Model(&tag).Update("title", cr.Title).Error
	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage("标签修改失败", c)
		return
	}
	res.OKWithMessage("标签修改成功", c)
}
