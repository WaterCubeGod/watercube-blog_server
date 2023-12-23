package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (r *RouterGroup) ArticleRouter() {
	articleApi := api.ApiGroupApp.ArticleApi
	r.POST("/article", middleware.JwtAuth(), articleApi.ArticleCreateView)
	r.PUT("/article_publish/:id", middleware.JwtAuth(), articleApi.ArticlePublishView)
	r.PUT("/article_update/:id", middleware.JwtAuth(), articleApi.ArticleUpdateView)
	r.GET("/my_articles/:id", middleware.JwtAuth(), articleApi.ArticleMyListView)
	r.GET("/articles", articleApi.ArticleListView)
	r.POST("/article_detail/:id", articleApi.ArticleDetailView)
}
