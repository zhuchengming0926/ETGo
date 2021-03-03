/**
 * @File: httpGetDemo.go
 * @Author: zhuchengming
 * @Description:get请求参数获取
 * @Date: 2021/2/26 17:30
 */

package v1

import (
	"ETGo/helper"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func HttpGetDemo(ctx *gin.Context) {
	//获取参数方式
	r := ctx.Request
	err := r.ParseForm()
	if err != nil {
		helper.Render.RenderJsonFail(ctx, err)
		return
	}

	uid := r.Form["uid"] //[]string 类型
	uidStr, _ := jsoniter.MarshalToString(uid)
	helper.Render.RenderJsonSucc(ctx, uidStr)
	return
}
/*
使用postman测试：
获取浏览器直接访问：http://localhost:8955/et/v1/variousHttp/getdemo?uid=888&uid=888
服务端输出 ：[888 888]
*/

/*
小结：r.Form是url.Values字典类型，r.Form[“id”]取到的是一个数组类型。
因为http.request在解析参数的时候会将同名的参数都放进同一个数组里。

因为r.Form包含了GET、POST参数，POST参数优先，那么想只获取GET参数怎么办？可以改进代码：
    query := r.URL.Query()
	uid := query["uid"][0]
	fmt.Println(uid)
。
原文链接：https://blog.csdn.net/guyan0319/article/details/84674627
*/