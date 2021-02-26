/**
 * @File: main.go
 * @Author: zhuchengming
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
	engine.LoadHTMLFiles("statics/view/index.html")
	engine.MaxMultipartMemory = 8 << 20 //8MB 设置最大的上传文件的大小

	//开启jsoniter的模糊模式
	extra.RegisterFuzzyDecoders()

	helper.Init(engine)
	defer helper.Clear()

	router.Http(engine)
	// 启动web server
	_ = engine.Run(conf.BasicConf.Server.Address)
}
