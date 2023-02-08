package v1

import (
	"evernote-client/global"
	"evernote-client/middleware"
	"evernote-client/model"
	"evernote-client/model/request"
	"evernote-client/model/response"
	"evernote-client/service"
	"evernote-client/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"strings"
	"time"
)

type User struct {
	UserName string `json:"username"`
}

type Auth struct {
	Data    User `json:"data"`
	IsLogin bool `json:"isLogin"`
}

type AuthLogin struct {
	Data User   `json:"data"`
	Msg  string `json:"msg"`
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var l request.Login
	_ = c.ShouldBindJSON(&l)
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if service.CheckTicket(l.Ticket, l.RandStr) {
		u := &model.EvnUser{Username: l.Username, Password: l.Password}
		if err, user := service.Login(u); err != nil {
			global.LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
		} else {
			tokenNext(c, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}
}

// @Tags Base
// @Summary 用户退出登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Logout(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := model.EvnJwtBlacklist{Jwt: token}
	if err := service.JsonInBlacklist(jwt); err != nil {
		global.LOG.Error("退出登录失败!", zap.Any("err", err))
		response.FailWithMessage("退出登录失败", c)
	} else {
		response.OkWithMessage("成功退出登录", c)
	}
}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user model.EvnUser) {
	// 多点登录
	if global.CONFIG.System.UseMultipoint {
		if err, jwtStr := service.GetRedisJWT(user.UUID.String()); err != redis.Nil {
			j := middleware.NewJWT()
			claims, err := j.ParseToken(jwtStr)
			if err == nil {
				response.OkWithDetailed(response.LoginResponse{
					User:      user,
					Token:     jwtStr,
					ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
				}, "登录成功", c)
				return
			}
		}
	}

	j := &middleware.JWT{SigningKey: []byte(global.CONFIG.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		UUID:       user.UUID,
		ID:         user.ID,
		Username:   user.Username,
		BufferTime: global.CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                          // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "qmtPlus",                                         // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("获取token失败!", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.CONFIG.System.UseMultipoint {
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	if err, jwtStr := service.GetRedisJWT(user.UUID.String()); err == redis.Nil {
		if err := service.SetRedisJWT(token, user.UUID.String()); err != nil {
			global.LOG.Error("设置登录状态失败!", zap.Any("err", err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.LOG.Error("设置登录状态失败!", zap.Any("err", err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT model.EvnJwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := service.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := service.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// @Summary 用户更新昵称
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/nickName [PATCH]
func UpdateNickName(c *gin.Context) {
	var user model.EvnUser
	err := c.ShouldBindJSON(&user)
	if err != nil || user.NickName == "" {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := service.UpdateNickName(getUserID(c), user.NickName); err != nil {
		global.LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败！"+err.Error(), c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Summary 用户修改密码
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [PATCH]
func ChangePassword(c *gin.Context) {
	var passWord request.ChangePassword
	err := c.ShouldBindJSON(&passWord)
	if err != nil || passWord.OldPass == "" || passWord.NewPass == "" {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := service.ChangePassword(getUserID(c), passWord); err != nil {
		global.LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags User
// @Summary 用户上传头像
// @Produce  application/json
// @Param file File
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /user/uploadAvatar [post]
func UploadAvatar(c *gin.Context) {
	var file model.EvnUpload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	err, file = service.UploadAvatar(getUserID(c), header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.LOG.Error("修改数据库链接失败!", zap.Any("err", err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(response.FileUploadResponse{File: file}, "上传成功", c)
}

// @Summary 用户修改邮箱
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/email [PATCH]
func UpdateEmail(c *gin.Context) {
	var user request.ChangeEmail
	err := c.ShouldBindJSON(&user)
	if err != nil || user.Email == "" || user.PassWord == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	redisKey := fmt.Sprintf("verify:%s", user.Email)
	err, value := service.GetRedis(redisKey)
	if err != nil || value != strings.TrimSpace(user.VerifyCode) {
		response.FailWithMessage("验证码错误", c)
		return
	}
	_ = service.DelRedis(redisKey)
	if err := service.UpdateEmail(getUserID(c), user); err != nil {
		global.LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败，密码与当前账户不符", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags SysUser
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body systemReq.Register true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {object} response.Response{data=systemRes.SysUserResponse,msg=string} "用户注册账号,返回包括用户信息"
// @Router /user/register [post]
func Register(c *gin.Context) {
	var r request.Register
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	redisKey := fmt.Sprintf("verify:%s", r.Email)
	err, value := service.GetRedis(redisKey)
	if err != nil || value != strings.TrimSpace(r.VerifyCode) {
		response.FailWithMessage("验证码错误", c)
		return
	}
	_ = service.DelRedis(redisKey)
	err = service.Register(r)
	if err != nil {
		global.LOG.Error("注册失败!", zap.Error(err))
		response.FailWithMessage("注册失败 "+err.Error(), c)
	} else {
		response.OkWithMessage("注册成功", c)
	}
}
