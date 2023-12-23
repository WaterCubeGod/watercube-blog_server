package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type TagRequest struct {
	Titles []string `json:"titles" binding:"required" msg:"请输入标签"` // 显示的标签
}

// TagCreateView 标签创建
// @Tags 标签管理
// @Summary 标签创建
// @Description 标签创建
// @Param data body TagRequest true "标签的title"
// @Router /api/tags [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (*TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, cr, "msg", c)
		return
	}
	num := 0
	tx := global.DB.Begin()
	for _, title := range cr.Titles {
		count := tx.Where("title = ?", title).Find(&models.TagModel{}).RowsAffected
		if count == 1 {
			continue
		}
		err := tx.Create(&models.TagModel{
			Title: title,
		}).Error
		if err != nil {
			res.FailWithMessage("标签创建失败", c)
			return
		}
		num++
	}
	tx.Commit()
	res.OKWithMessage(fmt.Sprintf("共创建%d个标签", num), c)
}
