package image_ser

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils"
	"io"
	"mime/multipart"
	"path"
	"strings"
)

const (
	B2M = 1024 * 1024 // 字节转换为MB
)

// WhiteImageList 图片上传的白名单
var WhiteImageList = []string{
	"jpg",
	"png",
	"jpeg",
	"ico",
	"tiff",
	"gif",
	"svg",
	"webp",
}

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

// ImageUploadService 处理文件上传的方法
func (*ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {

	fileName := file.Filename

	// 判断是否在白名单内
	fileList := strings.Split(fileName, ".")
	suffix := strings.ToLower(fileList[len(fileList)-1])
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = fmt.Sprintf("文件非法")
		return
	}
	//判断文件大小
	size := file.Size
	if size > int64(B2M*global.CONFIG.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过%dMB", global.CONFIG.Upload.Size)
		return
	}
	// 读取文件内容 hash
	filePath := path.Join(global.CONFIG.Upload.Path, file.Filename)
	res.FileName = filePath
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
		res.Msg = fmt.Sprintf("图片已存在")
		res.FileName = bannerModel.Path
		return
	}

	res.Msg = fmt.Sprintf("图片上传成功")
	res.IsSuccess = true

	// 图片入库
	global.DB.Create(&models.BannerModel{
		Path: filePath,
		Hash: imageHash,
		Name: fileName,
	})
	return
}
