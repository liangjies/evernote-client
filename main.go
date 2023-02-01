// evernote-client project main.go
package main

import (
	"evernote-client/core"
	"evernote-client/global"
	"evernote-client/initialize"
)

func main() {
	core.Viper()                  // 初始化Viper
	global.LOG = core.Zap()       // 初始化zap日志库
	global.DB = initialize.Gorm() // gorm连接数据库
	if global.DB != nil {
		initialize.MysqlTables(global.DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}

	core.RunServer()
}
