

/**
 * @File: conf.go
 * @Author: zhuchengming
 * @Description:
 * @Date: 2021/2/7 20:19
 */

package conf

import (
	"ETGo/components/base"
	"ETGo/elog"
	"ETGo/env"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

var (
	BasicConf TBasic
	Rconf ResourceConf
)

type TBasic struct {
	Log    elog.LogConfig
	Server env.ServerConfig
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
	path = filepath.Join(env.GetConfDirPath(), confType, filename)
	if yamlFile, err := ioutil.ReadFile(path); err != nil {
		panic(path + " get error: %v " + err.Error())
	} else if err = yaml.Unmarshal(yamlFile, s); err != nil {
		panic(path + " unmarshal error: %v" + err.Error())
	}
}

