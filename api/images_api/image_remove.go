package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// ImageRemoveView 删除图片
// @Tags 图片管理
// @Summary 删除图片
// @Description 删除图片
// @Param images body models.RemoveRequest true "删除图片的列表"
// @Router /api/images [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (*ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var imagesList []models.BannerModel
	count := global.DB.Find(&imagesList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}
	global.DB.Delete(&imagesList)
	res.OKWithMessage(fmt.Sprintf("共删除%d张图片", count), c)
}
