package article_api

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"strconv"
	"time"
)

func (*ArticleApi) ArticleDetailView(c *gin.Context) {
	strId := c.Param("id")
	// 查看文章是否发布
	isPublished, err := IsPublished(strId)
	if err != nil {
		global.LOG.Error(err)
		res.FailWithMessage("获取文章失败", c)
		return
	}
	if !isPublished {
		res.FailWithMessage("该文章还未发布", c)
		return
	}
	// 查看细节
	var article ctype.ArticleRequest
	// 缓存中没有
	if global.RDB.Get(PREFIX+strId+":title").Val() == "" {
		// 从es中获取
		resource, err := global.ES.Get().
			Index(global.ARTICLE_INDEX).
			Id(strId).
			Do(context.Background())
		if err != nil {
			global.LOG.Info(err)
			res.FailWithMessage("文章不存在", c)
			return
		}
		err = json.Unmarshal(resource.Source, &article)
		if err != nil {
			global.LOG.Error(err)
			res.FailWithMessage("文章获取失败", c)
			return
		}
		// 放入redis中
		ArticleRDB(&article, strId)
	} else {
		// 从redis中获取
		err := GetArticle(strId, &article)
		if err != nil {
			global.LOG.Error(err)
			res.FailWithMessage("获取文章内容失败", c)
			return
		}
	}
	res.OKWithData(article, c)
}

// GetArticle 从redis中获取数据
func GetArticle(strId string, article *ctype.ArticleRequest) error {
	prefix := PREFIX + strId
	article.Title = global.RDB.Get(prefix + ":title").Val()
	update, err := time.Parse("2006-01-02 15:04:05 +0800 CST", global.RDB.Get(prefix+":updated_at").Val())
	if err != nil {
		return err
	}
	article.UpdatedAt = update
	create, err := time.Parse("2006-01-02 15:04:05 +0800 CST", global.RDB.Get(prefix+":created_at").Val())
	if err != nil {
		return err
	}
	article.CreatedAt = create
	article.Abstract = global.RDB.Get(prefix + ":abstract").Val()
	article.Content = global.RDB.Get(prefix + ":content").Val()
	article.Category = global.RDB.Get(prefix + ":category").Val()
	article.Source = global.RDB.Get(prefix + ":source").Val()
	tags := global.RDB.Get(prefix + ":tags").Val()
	err = json.Unmarshal([]byte(tags), &article.Tags)
	if err != nil {
		return err
	}
	userID, err := strconv.Atoi(global.RDB.Get(prefix + ":user_id").Val())
	if err != nil {
		return err
	}
	article.UserID = uint(userID)
	article.Link = global.RDB.Get(prefix + ":link").Val()
	bannerID, err := strconv.Atoi(global.RDB.Get(prefix + ":banner_id").Val())
	if err != nil {
		return err
	}
	article.BannerID = uint(bannerID)
	article.BannerPath = global.RDB.Get(prefix + ":banner_path").Val()
	article.NickName = global.RDB.Get(prefix + ":nick_name").Val()
	look, err := strconv.Atoi(global.RDB.Get(prefix + ":look_count").Val())
	if err != nil {
		return err
	}
	article.LookCount = look
	comment, err := strconv.Atoi(global.RDB.Get(prefix + ":comment_count").Val())
	if err != nil {
		return err
	}
	article.CommentCount = comment
	digg, err := strconv.Atoi(global.RDB.Get(prefix + ":digg_count").Val())
	if err != nil {
		return err
	}
	article.DiggCount = digg
	collects, err := strconv.Atoi(global.RDB.Get(prefix + ":collects_count").Val())
	if err != nil {
		return err
	}
	article.CollectsCount = collects
	return nil
}
