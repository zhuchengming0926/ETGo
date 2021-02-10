/**************************************************************************
 *
 **************************************************************************/

/**
 * @File: logger.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description: 日志中间件
 * @Date: 2021/2/8 17:03
 */

package middleware

import (
	"ETGo/components/utils"
	"ETGo/elog"
	"ETGo/env"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	printRequestLen  = 10240
	printResponseLen = 10240
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func GetLogger() *logrus.Logger {
	logClient := logrus.New()

	//禁止logrus的输出
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	logClient.Out = src
	logClient.SetLevel(logrus.DebugLevel)
	apiLogPath := "./logs/access.log"

	logWriter, err := rotatelogs.New(
		apiLogPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(apiLogPath), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour), // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	logClient.AddHook(lfHook)
	return logClient
}


func Logger() gin.HandlerFunc {
	return func (c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 请求报文
		var requestBody []byte
		if c.Request.Body != nil {
			var err error
			requestBody, err = c.GetRawData()
			if err != nil {
				elog.Warnf(c, "get http request body error: %s", err.Error())
			}
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}

		blw := new(bodyLogWriter)
		if printResponseLen <= 0 {
			blw = &bodyLogWriter{body: nil, ResponseWriter: c.Writer}
		} else {
			blw = &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		}
		c.Writer = blw

		c.Set("handler", c.HandlerName())
		logID := elog.GetLogID(c)
		requestID := elog.GetRequestID(c)
		// 处理请求
		c.Next()

		response := ""
		if blw.body != nil {
			if len(blw.body.String()) <= printResponseLen {
				response = blw.body.String()
			} else {
				response = blw.body.String()[:printResponseLen]
			}
		}

		bodyStr := string(requestBody)
		if c.Request.URL.RawQuery != "" {
			bodyStr += "&" + c.Request.URL.RawQuery
		}
		if len(bodyStr) > printRequestLen {
			bodyStr = bodyStr[:printRequestLen]
		}

		// 结束时间
		end := time.Now()
		//执行时间
		clientIP := c.ClientIP()

		// 固定notice
		commonFields := []elog.Field{
			elog.String("logId", logID),
			elog.String("requestId", requestID),
			elog.String("localIp", utils.GetLocalIp()),
			elog.String("module", env.AppName),
			elog.String("cuid", getReqValueByKey(c, "cuid")),
			elog.String("device", getReqValueByKey(c, "device")),
			elog.String("channel", getReqValueByKey(c, "channel")),
			elog.String("os", getReqValueByKey(c, "os")),
			elog.String("vc", getReqValueByKey(c, "vc")),
			elog.String("vcname", getReqValueByKey(c, "vcname")),
			elog.String("userid", getReqValueByKey(c, "userid")),
			elog.String("uri", c.Request.RequestURI),
			elog.String("host", c.Request.Host),
			elog.String("method", c.Request.Method),
			elog.String("httpProto", c.Request.Proto),
			elog.String("handle", c.HandlerName()),
			elog.String("userAgent", c.Request.UserAgent()),
			elog.String("refer", c.Request.Referer()),
			elog.String("clientIp", clientIP),
			elog.String("cookie", getCookie(c)),
			elog.String("requestStartTime", utils.GetFormatRequestTime(start)),
			elog.String("requestEndTime", utils.GetFormatRequestTime(end)),
			elog.Float64("cost", utils.GetRequestCost(start, end)),
			elog.String("requestParam", bodyStr),
			elog.Int("responseStatus", c.Writer.Status()),
			elog.String("response", response),
		}

		elog.InfoLogger(c, "notice", commonFields...)

		//GetLogger().Infof("| %3d | %13v | %15s | %s  %s |",
		//	statusCode,
		//	latency,
		//	clientIP,
		//	method, path,
		//)
	}
}

// 从request body中解析特定字段作为notice key打印
func getReqValueByKey(ctx *gin.Context, k string) string {
	if vs, exist := ctx.Request.Form[k]; exist && len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func getCookie(ctx *gin.Context) string {
	cStr := ""
	for _, c := range ctx.Request.Cookies() {
		cStr += fmt.Sprintf("%s=%s&", c.Name, c.Value)
	}
	return strings.TrimRight(cStr, "&")
}