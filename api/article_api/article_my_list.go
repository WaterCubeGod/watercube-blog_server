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
	"strconv"
)

// ArticleMyListView 我的文章列表
// @Tags 文章管理
// @Summary 我的文章列表
// @Description 我的文章列表
// @Param key body models.PageInfo false "查询关键词"
// @Router /api/my_articles [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[ctype.ArticleRequest]}
func (*ArticleApi) ArticleMyListView(c *gin.Context) {
	var page models.PageInfo
	err := c.ShouldBindJSON(&page)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	from := (page.Page - 1) * page.Limit
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 获取查询条件
	keyFields := []string{"abstract", "title",
		"category", "tags", "nick_name"}
	query := elastic.NewBoolQuery().
		Must(elastic.NewTermQuery("user_id", id))
	if page.Key != "" {
		query = query.Must(elastic.NewMultiMatchQuery(page.Key, keyFields...))
	}
	// 启用高亮
	highlight := elastic.NewHighlight().
		Fields(*Str2High(keyFields)...).
		PreTags("<em>").PostTags("</em>")
	// 查找的字段
	fields := []string{"title", "updated_at", "banner_id", "banner_path", "abstract", "nick_name",
		"tags", "collects_count", "look_count", "comment_count", "digg_count", "is_publish", "category"}
	resource, err := global.ES.
		Search(global.ARTICLE_INDEX).
		Query(query).
		Sort("updated_at", true).
		From(from).
		Size(page.Limit).
		FetchSourceContext(elastic.NewFetchSourceContext(true).Include(fields...)).
		Highlight(highlight).
		Do(context.Background())
	count := resource.Hits.TotalHits.Value
	// 处理数据
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

func Str2High(highlightFields []string) *[]*elastic.HighlighterField {
	// 将字符串数组转换为 []*HighlighterField 切片
	var highlighterFields []*elastic.HighlighterField
	for _, field := range highlightFields {
		highlighterFields = append(highlighterFields, elastic.NewHighlighterField(field))
	}
	return &highlighterFields
}
