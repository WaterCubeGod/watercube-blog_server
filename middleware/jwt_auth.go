package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/utils/jwts"
)

// JwtAuth 用户权限中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		claims := JwtLogin(c, token)
		if claims == nil {
			return
		}
		// 判断是否在redis中
		key := global.RDB.Keys(fmt.Sprintf("logout_%s", token)).Val()
		if len(key) == 1 {
			res.FailWithMessage("您还未登录", c)
			c.Abort()
			return
		}
		// 登录的用户
		c.Set("claims", claims)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		claims := JwtLogin(c, token)
		if claims == nil {
			return
		}
		// 登录的用户
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMessage("没有权限", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

func JwtLogin(c *gin.Context, token string) *jwts.CustomClaims {
	if token == "" {
		res.FailWithMessage("未登录", c)
		c.Abort()
		return nil
	}
	claims, err := jwts.ParseToken(token)
	if err != nil {
		res.FailWithMessage("token错误", c)
		c.Abort()
		return nil
	}
	return claims
}
