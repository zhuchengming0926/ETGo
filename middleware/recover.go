

/**
 * @File: recover.go
 * @Author: zhuchengming
 * @Description:
 * @Date: 2021/2/8 17:07
 */

package middleware

import (
	"ETGo/helper"
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func Recover(ctx *gin.Context) {
	defer CatchRecoverRpc(ctx)
	ctx.Next()
}

// 针对rpc接口的处理
func CatchRecoverRpc(c *gin.Context) {
	// panic捕获
	if err := recover(); err != nil {
		// 请求url
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}
		// 请求报文
		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		helper.Render.RenderJsonAbort(c)
	}
}


