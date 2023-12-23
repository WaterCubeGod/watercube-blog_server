package ctype

import (
	"time"
)

type ArticleRequest struct {
	Title         string    `json:"title"`
	Abstract      string    `json:"abstract"`
	Content       string    `json:"content"`
	Category      string    `json:"category"`
	Source        string    `json:"source"`
	Tags          Array     `json:"tags"`
	UserID        uint      `json:"user_id"`
	Link          string    `json:"link"`
	BannerID      uint      `json:"banner_id"`
	BannerPath    string    `json:"banner_path"`
	NickName      string    `json:"nick_name"`
	LookCount     int       `json:"look_count"`    // 浏览量
	CommentCount  int       `json:"comment_count"` // 评论量
	DiggCount     int       `json:"digg_count"`    // 点赞量
	CollectsCount int       `json:"collects_count"`
	CreatedAt     time.Time `json:"created_at"` // 创建时间
	UpdatedAt     time.Time `json:"updated_at"` // 更新时间
	IsPublish     bool      `json:"is_publish"`
}

func (a *ArticleRequest) ApplyHighlight(field string, highlightResult []string) {
	switch field {
	case "content":
		a.Content = highlightResult[0]
	case "abstract":
		a.Abstract = highlightResult[0]
	case "title":
		a.Title = highlightResult[0]
	case "category":
		a.Category = highlightResult[0]
	case "tags":
		// 如果是 tags 字段，将高亮片段追加到 tags 数组中
		for _, tagHighlight := range highlightResult {
			a.Tags = append(a.Tags, tagHighlight)
		}
	case "nick_name":
		a.NickName = highlightResult[0]
	}
}
