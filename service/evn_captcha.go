package service

import (
	"crypto/tls"
	"encoding/json"
	"evernote-client/global"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/imroc/req/v3"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

//	type VerifyResult struct {
//		Data struct {
//			Retcode int `json:"retcode,omitempty"`
//		} `json:"data,omitempty"`
//		ReCode *int `json:"reCode,omitempty"`
//	}
type VerifyResult struct {
	Data   string `json:"data,omitempty"`
	ReCode *int   `json:"reCode,omitempty"`
}

func CheckTicket(ticket, randstr string) bool {
	if ticket == "" || randstr == "" {
		global.LOG.Error("验证码参数为空", zap.String("ticket", ticket), zap.String("randstr", randstr))
		return false
	}

	var r = rand.NewSource(time.Now().UnixNano()).Int63()
	var api = fmt.Sprintf("https://cgi.urlsec.qq.com/index.php?m=check&a=gw_check&url=https://www.qq.com/%d&randstr=%s&ticket=%s", r, randstr, ticket)

	global.LOG.Info("调用腾讯验证码API", zap.String("api", api))

	client := req.C().SetUserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36").DevMode()
	resp, err := client.R().SetHeaders(map[string]string{
		"Referer": "https://urlsec.qq.com/check.html",
	}).Get(api)
	if err != nil {
		global.LOG.Error("腾讯验证码API请求失败", zap.Error(err))
		return false
	}

	var data = resp.String()
	global.LOG.Info("腾讯验证码API响应", zap.String("response", data))

	data = strings.TrimLeft(data, "(")
	data = strings.TrimRight(data, ")")

	var res VerifyResult
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		global.LOG.Error("解析腾讯验证码响应失败", zap.Error(err), zap.String("data", data))
		return false
	}
	if res.ReCode == nil {
		global.LOG.Error("腾讯验证码响应中ReCode为空")
		return false
	}

	global.LOG.Info("腾讯验证码验证结果", zap.Int("reCode", *res.ReCode))
	if *res.ReCode == 0 {
		return true
	}
	return false
}

func SendVerifyCode(mail string) (err error) {
	randNum := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	if err = SetRedis(fmt.Sprintf("verify:%s", mail), randNum, 10*60); err != nil {
		return err
	}
	if err = Send(mail, randNum); err != nil {
		return err
	}
	return err
}

// Send 发送邮件
func Send(to string, code string) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", global.CONFIG.Mail.From, global.CONFIG.Mail.Nickname)
	m.SetHeader("To", to)
	m.SetHeader("Subject", fmt.Sprintf("用户注册验证码：%s - note.icewx.com", code))
	m.SetBody("text/html", fmt.Sprintf("您注册的验证码为：%s，10分钟内有效", code))
	mailer := gomail.NewDialer(global.CONFIG.Mail.Host, global.CONFIG.Mail.Port, global.CONFIG.Mail.From, global.CONFIG.Mail.Secret)
	mailer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err = mailer.DialAndSend(m); err != nil {
		global.LOG.Error(err.Error())
		return err
	}

	return nil
}
