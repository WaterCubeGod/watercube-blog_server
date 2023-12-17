package routers

import (
	"gvb_server/api"
)

func (r *RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	r.GET("/settings/:name", settingsApi.SettingsInfoView)
	r.PUT("/settings/:name", settingsApi.SettingsUpdateView)
}
