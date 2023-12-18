package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type ImageResponse struct {
	ID   uint   `json:"id"`                  // 图片id
	Path string `json:"path"`                // 图片路径
	Name string `gorm:"size:38" json:"name"` // 图片名称
}

// ImageNameListView 查看图片简单信息
// @Tags 图片管理
// @Summary 查看图片简单信息
// @Description 查看图片简单信息
// @Router /api/image_names [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[ImageResponse]}
func (*ImagesApi) ImageNameListView(c *gin.Context) {
	var imageList []ImageResponse

	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageList)

	res.OKWithData(imageList, c)
}
