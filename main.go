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
	"time"
)

type MateriasInfo struct {
	Id uint64 `json:"id"`
	MpId uint64 `json:"mp_id"`
	StoreId uint64 `json:"store_id"`
	CustomerId uint64 `json:"customer_id"`
	SkuSeriesId uint64 `json:"sku_series_id"`
	TemplateId uint64 `json:"template_id"`
	OrderId uint64 `json:"order_id"`
	SkuId uint64 `json:"sku_id"`
	TemplateType int `json:"template_type"`//模版类型 1 图片 2 视频插图片 3 视频插视频
	Version int `json:"version"`//商品最新版本
	Result string `json:"result"` //商品素材制作结果地址
	State int `json:"state"`//素材制作状态 0:已入库为发送至kafka；1:发送至kafka制作中；2:制作取消；4:制作完成;8:素材更新；16:更新完成；32:素材删除
IsDeleted int `json:"is_deleted"`
StartAt time.Time `json:"start_at"`
Timeout int `json:"timeout"`//素材制作超时时间
}

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
