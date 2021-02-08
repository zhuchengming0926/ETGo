/**************************************************************************
 * Copyright (c) 2020 Zuoyebang Inc. All Rights Reserved
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
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine)  {
	PreInit()
	InitResource(engine)
}

func PreInit()  {
	conf.SetAppName("ETGo")
	conf.InitConf()
}

func InitResource(engine *gin.Engine)  {
	InitMysql()
}