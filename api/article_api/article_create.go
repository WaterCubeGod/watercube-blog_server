package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"strconv"
	"time"
)

var PREFIX = "article:"

// ArticleCreateView 创建文章
// @Tags 文章管理
// @Summary 创建文章
// @Description 创建文章
// @Param data body ctype.ArticleRequest true "文章相关属性"
// @Router /api/article [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (*ArticleApi) ArticleCreateView(c *gin.Context) {
	var cr ctype.ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	tx := global.DB.Begin()
	// 创建文章
	count := tx.Where("title = ?", cr.Title).Find(&models.ArticleModel{}).RowsAffected
	if count > 0 {
		res.FailWithMessage("文章名称重复", c)
		return
	}
	err = tx.Create(&models.ArticleModel{
		Title:      cr.Title,
		UserID:     cr.UserID,
		BannerID:   cr.BannerID,
		BannerPath: cr.BannerPath,
	}).Error
	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage("创建文章失败", c)
		return
	}
	var article models.ArticleModel
	tx.Where("title = ?", cr.Title).Find(&article)
	// 存入es
	_, err = global.ES.Index().
		Index(global.ARTICLE_INDEX).
		Id(strconv.Itoa(int(article.ID))).
		BodyJson(&ctype.ArticleRequest{
			Title:         cr.Title,
			Abstract:      cr.Abstract,
			Content:       cr.Content,
			Category:      cr.Category,
			Source:        cr.Source,
			Link:          cr.Link,
			Tags:          cr.Tags,
			NickName:      cr.NickName,
			IsPublish:     false,
			LookCount:     0,
			CommentCount:  0,
			DiggCount:     0,
			CollectsCount: 0,
			CreatedAt:     article.CreatedAt,
			UpdatedAt:     article.UpdatedAt,
			BannerID:      cr.BannerID,
			BannerPath:    cr.BannerPath,
			UserID:        cr.UserID,
		}).Do(context.Background())
	if err != nil {
		res.FailWithMessage("创建文章失败", c)
		return
	}
	cr.CreatedAt = article.CreatedAt
	cr.UpdatedAt = article.UpdatedAt
	// 存入redis
	ArticleRDB(&cr, strconv.Itoa(int(article.ID)))
	tx.Commit()
	res.OKWithMessage("创建文章成功", c)
}

func ArticleRDB(article *ctype.ArticleRequest, id string) {
	prefix := PREFIX + id
	global.RDB.Set(prefix+":title", article.Title, time.Hour*4)
	global.RDB.Set(prefix+":updated_at", article.UpdatedAt.String(), time.Hour*4)
	global.RDB.Set(prefix+":created_at", article.CreatedAt.String(), time.Hour*4)
	global.RDB.Set(prefix+":abstract", article.Abstract, time.Hour*4)
	global.RDB.Set(prefix+":content", article.Content, time.Hour*4)
	global.RDB.Set(prefix+":category", article.Category, time.Hour*4)
	global.RDB.Set(prefix+":source", article.Source, time.Hour*4)
	tags, _ := json.Marshal(article.Tags)
	global.RDB.Set(prefix+":tags", string(tags), time.Hour*4)
	global.RDB.Set(prefix+":user_id", article.UserID, time.Hour*4)
	global.RDB.Set(prefix+":link", article.Link, time.Hour*4)
	global.RDB.Set(prefix+":banner_id", article.BannerID, time.Hour*4)
	global.RDB.Set(prefix+":banner_path", article.BannerPath, time.Hour*4)
	global.RDB.Set(prefix+":nick_name", article.NickName, time.Hour*4)
	global.RDB.Set(prefix+":look_count", article.LookCount, time.Hour*4)
	global.RDB.Set(prefix+":comment_count", article.CommentCount, time.Hour*4)
	global.RDB.Set(prefix+":digg_count", article.DiggCount, time.Hour*4)
	global.RDB.Set(prefix+":collects_count", article.CollectsCount, time.Hour*4)
}
