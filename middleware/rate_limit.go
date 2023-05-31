package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

/**
 * @Author: chengming1
 * @Date: 2023/5/31 16:34
 * @Desc: 接口限流
 */

/**
在做一个基于图片文字识别的题库管理系统，使用 Golang 调用百度 OCR 文字识别接口, 但是百度 OCR 接口有调用频率限制：
免费版的 QPS 为 2。即每秒最多调用两次
付费版的 QPS 为 10
如果不限速，就会报错：
原文链接：https://www.sunzhongwei.com/speed-limit-golang-gin-api
*/

const (
	LIMIT_PER_SECOND = 1 // 限速值
	BUCKET_SIZE      = 1 // burst代表桶的容量大小
)

var limiter *rate.Limiter

func init() {
	limiter = rate.NewLimiter(LIMIT_PER_SECOND, BUCKET_SIZE)
}

func OcrRequestLimit(ctx *gin.Context) {
	limiter.Wait(ctx) // 阻塞式等待，不丢弃请求，可以设置ctx的deadline或者timeout，决定wait的最长时间
	ctx.Next()
}
