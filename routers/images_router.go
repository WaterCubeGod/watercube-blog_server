package routers

import "gvb_server/api"

func (r *RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	r.POST("images", imagesApi.ImageUploadView)
}
