package service

import (
	"encoding/json"
	"evernote-client/global"
	"fmt"
	"github.com/imroc/req/v3"
	"gopkg.in/gomail.v2"
	"math/rand"
	"strings"
	"time"
)

type VerifyResult struct {
	Data struct {
		Retcode int `json:"retcode,omitempty"`
	} `json:"data,omitempty"`
	ReCode *int `json:"reCode,omitempty"`
}

func CheckTicket(ticket, randstr string) bool {
	var r = rand.NewSource(time.Now().UnixNano()).Int63()
	var api = fmt.Sprintf("https://cgi.urlsec.qq.com/index.php?m=check&a=gw_check&url=https://www.qq.com/%d&randstr=%s&ticket=%s", r, randstr, ticket)
	client := req.C().SetUserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36").DevMode()
	resp, err := client.R().SetHeaders(map[string]string{
		"Referer": "https://urlsec.qq.com/check.html",
	}).Get(api)
	if err != nil {
		return false
	}
	var data = resp.String()
	data = strings.TrimLeft(data, "(")
	data = strings.TrimRight(data, ")")

	var res VerifyResult
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(*res.ReCode)
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
	if err = mailer.DialAndSend(m); err != nil {
		global.LOG.Error(err.Error())
		return err
	}

	return nil
}
