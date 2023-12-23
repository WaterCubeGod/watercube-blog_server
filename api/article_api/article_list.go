package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
)

// ArticleListView 文章列表
// @Tags 文章管理
// @Summary 文章列表
// @Description 文章列表
// @Param key body models.PageInfo false "查询关键词"
// @Router /api/articles [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[ArticleRequest]}
func (*ArticleApi) ArticleListView(c *gin.Context) {
	var page models.PageInfo
	err := c.ShouldBindJSON(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	from := (page.Page - 1) * page.Limit
	// 获取查询条件
	keyFields := []string{"abstract", "title",
		"category", "tags", "nick_name"}
	query := elastic.NewBoolQuery().
		Must(elastic.NewTermQuery("is_publish", true))
	if page.Key != "" {
		query = query.Must(elastic.NewMultiMatchQuery(page.Key, keyFields...))
	}
	// 启用高亮
	highlight := elastic.NewHighlight().
		Fields(*Str2High(keyFields)...).
		PreTags("<em>").PostTags("</em>")
	// 返回的字段
	fields := []string{"title", "updated_at", "banner_id", "banner_path", "abstract",
		"tags", "collects_count", "look_count", "comment_count", "digg_count", "category", "user_id",
		"nick_name", "is_publish"}
	// 查找
	resource, err := global.ES.
		Search(global.ARTICLE_INDEX).
		Query(query).
		Sort("updated_at", true).
		From(from).
		Size(page.Limit).
		Highlight(highlight).
		FetchSourceContext(elastic.NewFetchSourceContext(true).Include(fields...)).
		Do(context.Background())
	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage("查找失败", c)
		return
	}
	count := resource.Hits.TotalHits.Value
	var articles []ctype.ArticleRequest
	for _, hit := range resource.Hits.Hits {
		var article ctype.ArticleRequest
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.LOG.Error(err)
			res.FailWithMessage("查看列表失败", c)
			return
		}
		// 处理高亮
		hl := hit.Highlight
		for _, field := range keyFields {
			highlightResult, found := hl[field]
			if found {
				// 如果存在高亮结果，将其应用到原始文本中
				article.ApplyHighlight(field, highlightResult)
			}
		}
		articles = append(articles, article)
	}
	res.OKWithList(articles, count, c)
}
