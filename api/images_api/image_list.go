package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
)

// ImageListView 图片列表查询页
func (*ImagesApi) ImageListView(c *gin.Context) {
	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	fmt.Println(page)
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
