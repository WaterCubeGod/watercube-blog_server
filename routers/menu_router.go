package routers

import (
	"gvb_server/api"
)

func (r *RouterGroup) MenuRouter() {
	menuApi := api.ApiGroupApp.MenuApi
	r.GET("menus", menuApi.MenuListView)
	r.GET("menu_names", menuApi.MenuNameListView)
	r.POST("menus", menuApi.MenuCreateView)
	r.PUT("menus/:id", menuApi.MenuUpdateView)
	r.DELETE("menus", menuApi.MenuRemoveView)
	r.GET("menu_detail/:id", menuApi.MenuDetailView)
}
