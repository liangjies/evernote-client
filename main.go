// evernote-client project main.go
package main

import (
	"evernote-client/core"
	"evernote-client/global"
	"evernote-client/initialize"
)

func main() {
	core.Viper()                      // 初始化Viper
	global.SYS_LOG = core.Zap()       // 初始化zap日志库
	global.SYS_DB = initialize.Gorm() // gorm连接数据库
	if global.SYS_DB != nil {
		initialize.MysqlTables(global.SYS_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.SYS_DB.DB()
		defer db.Close()
	}

	core.RunServer()
}
