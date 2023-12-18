package routers

import "gvb_server/api"

func (r *RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	r.GET("images", imagesApi.ImageListView)
	r.GET("image_names", imagesApi.ImageNameListView)
	r.POST("images", imagesApi.ImageUploadView)
	r.DELETE("images", imagesApi.ImageRemoveView)
	r.PUT("images", imagesApi.ImageUpdateView)
}
