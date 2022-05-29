package core

import (
	"evernote-client/global"
	"evernote-client/initialize"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	Router := initialize.Routers()
	if global.SYS_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	//Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.SYS_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.SYS_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
		欢迎使用 Evernote-Client
		运行地址:http://127.0.0.1%s
		`, address)
	global.SYS_LOG.Error(s.ListenAndServe().Error())
}
