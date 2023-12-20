package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/plugins/email"
	"gvb_server/utils/jwts"
	"gvb_server/utils/random"
	"time"
)

type BindEmailRequest struct {
	Email string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code  *string `json:"code"`
}

// UserBindEmailView 用户绑定邮箱
// @Tags 用户管理
// @Summary 用户绑定邮箱
// @Description 用户绑定邮箱
// @Param data body BindEmailRequest true "邮箱绑定相关参数"
// @Router /api/user_bind_email [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (*UserApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	token := c.Request.Header.Get("token")
	// 用户绑定邮箱，第一次输入邮箱
	// 后台会给验证码，用户输入验证码，密码，完成绑定
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, "msg", c)
		return
	}
	if cr.Code == nil {
		// 第一次
		code := random.Code(4)
		emailApi := email.NewCode()
		// 将验证码存入redis
		global.RDB.Set(fmt.Sprintf("%s_code", token), code, time.Minute)
		err := emailApi.SendEmail(cr.Email, "你的验证码是 "+code)
		if err != nil {
			res.FailWithMessage("验证码发送失败", c)
			return
		}
		res.OKWithMessage("验证码已发送", c)
		return
	} else {
		code := global.RDB.Get(fmt.Sprintf("%s_code", token))
		if code.Val() != *cr.Code {
			res.OKWithMessage("验证码错误", c)
			return
		}
		var user models.UserModel
		err = global.DB.Take(&user, claims.UserID).Error
		if err != nil {
			res.FailWithMessage("用户不存在", c)
			return
		}
		err = global.DB.Model(&user).Update("email", cr.Email).Error
		if err != nil {
			res.FailWithMessage("邮箱绑定失败", c)
			return
		}
	}
	res.OKWithMessage("邮箱绑定成功", c)
}
