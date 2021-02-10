/**************************************************************************
 *
 **************************************************************************/

/**
 * @File: env.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/7 20:22
 */

package env

import (
	"log"
	"path/filepath"
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

func SetConfDirPath(subPath ...string) {
	confDirName = filepath.Join(subPath...)
	println("load conf: ", confDirName)
}

// GetConfDirPath 返回配置文件目录绝对地址
func GetConfDirPath() string {
	return confDirName
}

// LogRootPath 返回log目录的绝对地址
func GetLogDirPath() string {
	return filepath.Join(GetRootPath(), "logs")
}

// RootPath 返回应用的根目录
func GetRootPath() string {
	return rootPath
}



