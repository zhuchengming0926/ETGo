/**
 * @File: zap_logger.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/10 11:04
 */

package elog

import (
	"ETGo/components/utils"
	"ETGo/env"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetZapLogger() (l *zap.Logger) {
	if ZapLogger == nil {
		ZapLogger = newLogger().WithOptions(zap.AddCallerSkip(1))
	}
	return ZapLogger
}

func zapLogger(ctx *gin.Context) *zap.Logger {
	m := GetZapLogger()
	//m = m.WithOptions(zap.AddCallerSkip(1))
	if ctx == nil {
		return m
	}
	return m.With(
		zap.String("logId", GetLogID(ctx)),
		zap.String("requestId", GetRequestID(ctx)),
		zap.String("module", env.AppName),
		zap.String("localIp", utils.GetLocalIp()),
	)
}

func DebugLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if NoLog(ctx) {
		return
	}
	zapLogger(ctx).Debug(msg, fields...)
}
func InfoLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if NoLog(ctx) {
		return
	}
	zapLogger(ctx).Info(msg, fields...)
}

func WarnLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if NoLog(ctx) {
		return
	}
	zapLogger(ctx).Warn(msg, fields...)
}

func ErrorLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if NoLog(ctx) {
		return
	}
	zapLogger(ctx).Error(msg, fields...)
}

func PanicLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if NoLog(ctx) {
		return
	}
	zapLogger(ctx).Panic(msg, fields...)
}

func FatalLogger(ctx *gin.Context, msg string, fields ...zap.Field) {
	if NoLog(ctx) {
		return
	}
	zapLogger(ctx).Fatal(msg, fields...)
}
