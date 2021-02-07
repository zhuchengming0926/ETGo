/**************************************************************************
 * Copyright (c) 2020 Zuoyebang Inc. All Rights Reserved
 **************************************************************************/

/**
 * @File: conf.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/7 20:19
 */

package conf

import (
	"ETGo/components/base"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

var (
	BasicConf TBasic
	Rconf ResourceConf
)

type TBasic struct {
	Server ServerConfig
}

type ResourceConf struct {
	Mysql   map[string]base.MysqlConf
}


const (
	SubConfDefault = ""
	SubConfApp     = "app"
)

func InitConf()  {
	// 加载资源类配置（optional）
	LoadConf("resource.yaml", SubConfDefault, &Rconf)
	LoadConf("config.yaml", SubConfDefault, &BasicConf)
}

func LoadConf(filename, confType string, s interface{})  {
	var path string
	path = filepath.Join(GetConfDirPath(), confType, filename)
	if yamlFile, err := ioutil.ReadFile(path); err != nil {
		panic(filename + " get error: %v " + err.Error())
	} else if err = yaml.Unmarshal(yamlFile, s); err != nil {
		panic(filename + " unmarshal error: %v" + err.Error())
	}
}

func SetConfDirPath(subPath ...string) {
	confDirName = filepath.Join(subPath...)
	println("load conf: ", confDirName)
}

// GetConfDirPath 返回配置文件目录绝对地址
func GetConfDirPath() string {
	return confDirName
}