/**************************************************************************
 *
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
	"ETGo/router"
	"github.com/json-iterator/go/extra"
)

func main()  {
	engine := components.GetEngin()

	//开启jsoniter的模糊模式
	extra.RegisterFuzzyDecoders()

	helper.Init(engine)
	defer helper.Clear()

	router.Http(engine)
	// 启动web server
	_ = engine.Run(conf.BasicConf.Server.Address)
}
