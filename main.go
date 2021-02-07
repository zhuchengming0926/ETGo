/**************************************************************************
 * Copyright (c) 2020 Zuoyebang Inc. All Rights Reserved
 **************************************************************************/

/**
 * @File: main.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/7 19:49
 */

package main

import (
	"ETGo/components"
	"ETGo/conf"
	"ETGo/helper"
)

func main()  {
	engine := components.GetEngin()
	helper.Init(engine)

	// 启动web server
	_ = engine.Run(conf.BasicConf.Server.Address)
}
