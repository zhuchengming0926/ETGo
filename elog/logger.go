/**************************************************************************
 *
 **************************************************************************/

/**
 * @File: logger.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/9 16:39
 */

package elog

import (
	"ETGo/env"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// log文件后缀类型
const (
	txtLogNormal    = "normal"
	txtLogWarnFatal = "warnfatal"
	txtLogStdout    = "stdout"
)

type Field = zap.Field

var (
	Binary = zap.Binary
	Bool   = zap.Bool

	ByteString = zap.ByteString
	String     = zap.String
	Strings    = zap.Strings

	Float64 = zap.Float64
	Float32 = zap.Float32

	Int   = zap.Int
	Int64 = zap.Int64
	Int32 = zap.Int32
	Int16 = zap.Int16
	Int8  = zap.Int8

	Uint   = zap.Uint
	Uint64 = zap.Uint64
	Uint32 = zap.Uint32

	Reflect   = zap.Reflect
	Namespace = zap.Namespace
	Duration  = zap.Duration
	Object    = zap.Object
	Any       = zap.Any
	Skip      = zap.Skip()
)
var (
	SugaredLogger *zap.SugaredLogger
	ZapLogger     *zap.Logger
)

// NewLogger 新建Logger，每一次新建会同时创建x.log与x.log.wf (access.log 不会生成wf)
func newLogger() *zap.Logger {
	var infoLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logConfig.ZapLevel && lvl <= zapcore.InfoLevel
	})

	var errorLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logConfig.ZapLevel && lvl >= zapcore.WarnLevel
	})

	var stdLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= logConfig.ZapLevel && lvl >= zapcore.DebugLevel
	})

	name := env.AppName
	if name == "" {
		name = "server"
	}
	var zapCore []zapcore.Core
	if logConfig.Stdout {
		c := zapcore.NewCore(
			getEncoder(),
			zapcore.AddSync(getLogWriter(name, txtLogStdout)),
			stdLevel)
		zapCore = append(zapCore, c)
	}

	// 仅开发环境有效，便于开发调试
	if logConfig.Log2File {
		zapCore = append(zapCore,
			zapcore.NewCore(
				getEncoder(),
				zapcore.AddSync(getLogWriter(name, txtLogNormal)),
				infoLevel))

		zapCore = append(zapCore,
			zapcore.NewCore(
				getEncoder(),
				zapcore.AddSync(getLogWriter(name, txtLogWarnFatal)),
				errorLevel))
	}

	// core
	core := zapcore.NewTee(zapCore...)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()

	// 由于之前没有DPanic，同化DPanic和Panic
	development := zap.Development()

	// 设置初始化字段
	filed := zap.Fields()

	// 构造日志
	logger := zap.New(core, filed, caller, development)

	return logger
}

func getEncoder() zapcore.Encoder {
	// 公用编码器
	timeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}

	encoderCfg := zapcore.EncoderConfig{
		LevelKey:       "level",
		TimeKey:        "time",
		CallerKey:      "file",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}
	return NewETJSONEncoder(encoderCfg)
}

func getLogWriter(name, loggerType string) zapcore.WriteSyncer {
	// stdOut
	if loggerType == txtLogStdout {
		return zapcore.AddSync(os.Stdout)
	}

	// 打印到name.log 中
	logDir := logConfig.Path
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0777)
		if err != nil {
			panic(fmt.Errorf("log conf err: create log dir '%s' error: %s", logDir, err))
		}
	}

	filename := filepath.Join(strings.TrimSuffix(logDir, "/"), appendLogFileTail(name, loggerType))
	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic("open log file error: " + err.Error())
	}
	return zapcore.AddSync(fd)
}

// genFilename 拼装完整文件名
func appendLogFileTail(appName, loggerType string) string {
	var tailFixed string
	switch loggerType {
	case txtLogNormal:
		tailFixed = ".log"
	case txtLogWarnFatal:
		tailFixed = ".log.wf"
	default:
		tailFixed = ".log"
	}
	return appName + tailFixed
}

func CloseLogger() {
	if SugaredLogger != nil {
		_ = SugaredLogger.Sync()
	}

	if ZapLogger != nil {
		_ = ZapLogger.Sync()
	}
}