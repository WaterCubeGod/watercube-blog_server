package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
)

// ArticlePublishView 文章发布
// @Tags 文章管理
// @Summary 文章发布
// @Description 文章发布
// @Param id query int true "要发布的文章id"
// @Router /api/article_publish/:id [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (*ArticleApi) ArticlePublishView(c *gin.Context) {
	id := c.Param("id")
	// 查看文章是否发布
	isPublished, err := IsPublished(id)
	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage("获取文章失败", c)
		return
	}
	if isPublished {
		res.FailWithMessage("文章已经发布", c)
		return
	}
	// 更新为发布状态
	_, err = global.ES.Update().
		Index(global.ARTICLE_INDEX).
		Id(id).
		Doc(map[string]bool{
			"is_publish": true,
		}).
		DocAsUpsert(false).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage("文章不存在", c)
		return
	}
	res.OKWithMessage("文章发布成功", c)
}

func IsPublished(id string) (bool, error) {
	// 根据id获取文章状态
	resource, err := global.ES.Get().
		Index(global.ARTICLE_INDEX).
		Id(id).
		FetchSourceContext(elastic.NewFetchSourceContext(true).Include("is_publish")).
		Do(context.Background())
	if err != nil {
		return false, err
	}
	var article ctype.ArticleRequest
	err = json.Unmarshal(resource.Source, &article)
	if err != nil {
		return false, err
	}
	// 查看文章是否发布
	if article.IsPublish {
		return true, nil
	} else {
		return false, nil
	}
}
