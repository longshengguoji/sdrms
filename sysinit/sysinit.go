package sysinit

import (
	"github.com/astaxie/beego"
	"sdrms/utils"
)

func init() {
	//启用Session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//初始化日志
	utils.InitLogs()
	//初始化数据库
	InitDatabase()
}
