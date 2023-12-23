package article_api

import (
	"context"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"strconv"
)

// ArticleUpdateView 文章修改
// @Tags 文章管理
// @Summary 文章修改
// @Description 文章修改
// @Param data body ctype.ArticleRequest true "文章相关属性"
// @Router /api/article_update [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (*ArticleApi) ArticleUpdateView(c *gin.Context) {
	var cr ctype.ArticleRequest
	err := c.ShouldBindJSON(&cr)
	id := c.Param("id")
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	tx := global.DB.Begin()
	err = tx.Model(&models.ArticleModel{}).Where("id = ?", id).Updates(map[string]any{
		"title":       cr.Title,
		"banner_id":   cr.BannerID,
		"banner_path": cr.BannerPath,
	}).Error
	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage("文章修改错误", c)
		return
	}
	var article models.ArticleModel
	tx.Where("title = ?", cr.Title).Find(&article)
	// 存入redis
	ArticleRDB(&cr, strconv.Itoa(int(article.ID)))
	// 存入es
	_, err = global.ES.Update().
		Index(global.ARTICLE_INDEX).
		Id(strconv.Itoa(int(article.ID))).
		Id(strconv.Itoa(int(article.ID))).
		Doc(map[string]any{
			"title":      cr.Title,
			"abstract":   cr.Abstract,
			"content":    cr.Content,
			"category":   cr.Category,
			"source":     cr.Source,
			"tags":       cr.Tags,
			"updated_at": article.UpdatedAt,
		}).
		DocAsUpsert(false).
		Do(context.Background())
	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage("更新文章失败,文章可能不存在", c)
		return
	}
	tx.Commit()
	res.OKWithMessage("文章更新成功", c)
}
