package routers

import "gvb_server/api/menu_api"

func (r *RouterGroup) MenuRouter() {
	menuApi := menu_api.MenuApi{}
	r.GET("menus", menuApi.MenuListView)
	r.GET("menu_names", menuApi.MenuNameListView)
	r.POST("menus", menuApi.MenuCreateView)
}
