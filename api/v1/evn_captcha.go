package v1

import (
	"evernote-client/global"
	"evernote-client/model/request"
	"evernote-client/model/response"
	"evernote-client/service"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

// Captcha
// @Tags Base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"验证码获取成功"}"
// @Router /base/captcha [post]
func Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.CONFIG.Captcha.ImgHeight, global.CONFIG.Captcha.ImgWidth, global.CONFIG.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.LOG.Error("验证码获取失败!", zap.Any("err", err))
		response.FailWithMessage("验证码获取失败", c)
	} else {
		response.OkWithDetailed(response.SysCaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}

func SendVerifyCode(c *gin.Context) {
	var l request.SendVerifyCodeReq
	_ = c.ShouldBindJSON(&l)
	if service.CheckTicket(l.Ticket, l.RandStr) {
		if err := service.SendVerifyCode(l.Email); err != nil {
			global.LOG.Error("验证码发送失败!", zap.Any("err", err))
			response.FailWithMessage("验证码发送失败", c)
		} else {
			response.OkWithMessage("验证码发送成功", c)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}
}
