/**************************************************************************
 *
 **************************************************************************/

/**
 * @File: Init.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description: 初始化的一些操作
 * @Date: 2021/2/7 20:17
 */

package helper

import (
	"ETGo/conf"
	"ETGo/elog"
	"ETGo/env"
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine)  {
	env.SetAppName("ETGo")
	PreInit()
	InitResource(engine)
}

func PreInit()  {
	conf.InitConf()
	elog.InitLog(conf.BasicConf.Log)
}

func InitResource(engine *gin.Engine)  {
	InitMysql()
}

func Clear() {
	// 服务结束时的清理工作，对应 Init() 初始化的资源
	elog.CloseLogger()
}