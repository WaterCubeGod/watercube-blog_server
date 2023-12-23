package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (r *RouterGroup) TagRouter() {
	tagApi := api.ApiGroupApp.TagApi
	r.POST("/tags", middleware.JwtAuth(), tagApi.TagCreateView)
	r.GET("/tags", middleware.JwtAuth(), tagApi.TagListView)
	r.PUT("/tag", middleware.JwtAdmin(), tagApi.TagUpdateView)
	r.DELETE("/tags", middleware.JwtAdmin(), tagApi.TagRemoveView)
}
