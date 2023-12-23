package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

// ImageListView 图片列表查询
// @Tags 图片管理
// @Summary 图片列表查询
// @Description 图片列表查询
// @Param key body models.PageInfo false "文章列表"
// @Router /api/images [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.BannerModel]}
func (*ImagesApi) ImageListView(c *gin.Context) {
	var page models.PageInfo
	err := c.ShouldBindJSON(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, count, err := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: page,
		Debug:    true,
	})

	res.OKWithList(list, count, c)
}
