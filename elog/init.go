/**************************************************************************
 *
 **************************************************************************/

/**
 * @File: init.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/9 16:25
 */

package elog

import (
	"ETGo/env"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

// 对用户暴露的log配置
type LogConfig struct {
	Level  string `yaml:"level"`
	Stdout bool   `yaml:"stdout"`
}

type loggerConfig struct {
	ZapLevel zapcore.Level

	// 以下变量仅对开发环境生效
	Stdout   bool
	Log2File bool
	Path     string
}

// 全局配置 仅限Init函数进行变更
var logConfig = loggerConfig{
	ZapLevel: zapcore.InfoLevel,

	Stdout:   false,
	Log2File: true,
	Path:     "./logs",
}

func InitLog(config LogConfig) *zap.SugaredLogger {
	if err := RegisterETJSONEncoder(); err != nil {
		panic(err)
	}

	logConfig.ZapLevel = getLogLevel(config.Level)
	logConfig.Log2File = true
	logConfig.Stdout = config.Stdout
	logConfig.Path = env.GetLogDirPath()

	SugaredLogger = GetLogger()
	tt, _ := jsoniter.MarshalToString(logConfig)
	fmt.Println(tt)
	return SugaredLogger
}

func getLogLevel(lv string) (level zapcore.Level) {
	str := strings.ToUpper(lv)
	switch str {
	case "DEBUG":
		level = zap.DebugLevel
	case "INFO":
		level = zap.InfoLevel
	case "WARN":
		level = zap.WarnLevel
	case "ERROR":
		level = zap.ErrorLevel
	case "FATAL":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	return level
}