package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"path"
)

const (
	B2M = 1024 * 1024 // 字节转换为MB
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

// ImageUploadView 上传图片，返回图片的url
func (*ImagesApi) ImageUploadView(c *gin.Context) {
	multipartForm, err := c.MultipartForm()
	if err != nil {
		res.FailWithCode(res.UploadError, c)
		return
	}
	var resList []FileUploadResponse
	// 图片列表
	fileHeaders, ok := multipartForm.File["images"]
	if !ok {
		res.FailWithCode(res.UploadError, c)
		return
	}

	for _, file := range fileHeaders {
		size := file.Size
		if size > int64(B2M*global.CONFIG.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过%dMB", global.CONFIG.Upload.Size),
			})
			continue
		}
		filePath := path.Join(global.CONFIG.Upload.Path, file.Filename)
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			resList = append(resList, FileUploadResponse{
				FileName:  filePath,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片上传失败，存储错误"),
			})
			global.LOG.Error(err)
			continue
		}
		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       fmt.Sprintf("图片上传成功"),
		})
	}
	res.OKWithData(resList, c)
}
