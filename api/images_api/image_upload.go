package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/service/image_ser"
)

// ImageUploadView 上传图片
// @Tags 图片管理
// @Summary 上传图片
// @Description 上传图片
// @Param images formData string true "上传的一系列图片(实际参数:multipart.Form)"
// @Router /api/images [post]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[image_ser.FileUploadResponse]}
func (*ImagesApi) ImageUploadView(c *gin.Context) {
	multipartForm, err := c.MultipartForm()
	if err != nil {
		res.FailWithCode(res.UploadError, c)
		return
	}
	// 图片列表
	fileHeaders, ok := multipartForm.File["images"]
	if !ok {
		res.FailWithCode(res.UploadError, c)
		return
	}
	// 返回图片列表
	var resList []image_ser.FileUploadResponse

	for _, file := range fileHeaders {
		// 保存上传
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		err = c.SaveUploadedFile(file, serviceRes.FileName)
		if err != nil {
			global.LOG.Error(err)
			serviceRes.Msg = fmt.Sprintf("图片上传失败，存储错误")
			resList = append(resList, serviceRes)
			continue
		}
		resList = append(resList, serviceRes)
	}
	res.OKWithData(resList, c)
}
