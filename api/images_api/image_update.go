package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type ImageUpdate struct {
	ID   uint   `json:"id" binding:"required" msg:"请输入文件id"`
	Name string `json:"name" binding:"required" msg:"请输入文件名称"`
}

// ImageUpdateView 修改图片名称
// @Tags 图片管理
// @Summary 修改图片名称
// @Description 修改图片名称
// @Param images body ImageUpdate true "修改图片名称"
// @Router /api/images [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (*ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr ImageUpdate
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, "msg", c)
		return
	}
	var imageModel models.BannerModel
	err = global.DB.Take(&imageModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("文件不存在", c)
		return
	}
	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OKWithMessage("图片名称修改成功", c)
}
