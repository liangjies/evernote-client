package service

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
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
