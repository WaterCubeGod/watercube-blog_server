package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (r *RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	r.GET("/users", middleware.JwtAuth(), userApi.UserListView)
	r.POST("/email_login", userApi.EmailLoginView)
	r.PUT("/user_role", middleware.JwtAdmin(), userApi.UserUpdateRoleView)
	r.PUT("/user_pwd", middleware.JwtAuth(), userApi.UserUpdatePwd)
	r.POST("/logout", middleware.JwtAuth(), userApi.LogoutView)
	r.DELETE("/users", middleware.JwtAdmin(), userApi.UserRemoveView)
	r.PUT("/user_bind_email", middleware.JwtAuth(), userApi.UserBindEmailView)
}
