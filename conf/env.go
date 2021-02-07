/**************************************************************************
 * Copyright (c) 2020 Zuoyebang Inc. All Rights Reserved
 **************************************************************************/

/**
 * @File: env.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/7 20:22
 */

package conf

import (
	"log"
	"time"
)

var (
	AppName string

	// app root path
	rootPath    string
	confDirName string
)

type ServerConfig struct {
	Address      string        `yaml:"address"`
	ReadTimeout  time.Duration `yaml:"readtimeout"`
	WriteTimeout time.Duration `yaml:"writetimeout"`
}

func SetAppName(appName string)  {
	if appName == "" {
		panic("请使用 env.SetAppName(模块名) 指定appName，一旦创建不能再修改")
	}
	AppName = appName
	log.Println("appName is ", AppName)
	setAppPath()
}

func setAppPath() {
	SetRootPath("./")
	SetConfDirPath(GetRootPath(), "conf")
}

// SetRootPath 设置应用的根目录
func SetRootPath(r string) {
	rootPath = r
}

// RootPath 返回应用的根目录
func GetRootPath() string {
	return rootPath
}



