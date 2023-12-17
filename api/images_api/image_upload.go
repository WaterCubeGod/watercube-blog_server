package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils"
	"io"
	"path"
	"strings"
	"time"
)

const (
	B2M = 1024 * 1024 // 字节转换为MB
)

var (
	// WhiteImageList 图片上传的白名单
	WhiteImageList = []string{
		"jpg",
		"png",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"webp",
	}
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
		fileName := file.Filename
		// 判断是否在白名单内
		fileList := strings.Split(fileName, ".")
		suffix := strings.ToLower(fileList[len(fileList)-1])
		if !utils.InList(suffix, WhiteImageList) {
			resList = append(resList, FileUploadResponse{
				FileName:  fileName,
				IsSuccess: false,
				Msg:       fmt.Sprintf("文件非法"),
			})
			continue
		}
		size := file.Size
		if size > int64(B2M*global.CONFIG.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  fileName,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过%dMB", global.CONFIG.Upload.Size),
			})
			continue
		}

		filePath := path.Join(global.CONFIG.Upload.Path, file.Filename)
		fileObj, err := file.Open()
		if err != nil {
			global.LOG.Error(err)
		}
		byteData, err := io.ReadAll(fileObj)
		imageHash := utils.Md5(byteData)
		// 去数据库中查看文件是否存在（根据hash值）
		var bannerModel models.BannerModel
		err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
		if err == nil {
			// 找到了
			resList = append(resList, FileUploadResponse{
				FileName:  bannerModel.Path,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片已存在"),
			})
			continue
		}
		err = c.SaveUploadedFile(file, filePath)
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

		t := time.Now()
		m := models.MODEL{
			CreateAt: t,
			UpdateAt: t,
		}
		// 图片入库
		global.DB.Create(&models.BannerModel{
			MODEL: m,
			Path:  filePath,
			Hash:  imageHash,
			Name:  fileName,
		})
	}
	res.OKWithData(resList, c)
}
