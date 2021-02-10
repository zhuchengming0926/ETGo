

/**
 * @File: mysql.go
 * @Author: zhuchengming
 * @Description:
 * @Date: 2021/2/7 20:36
 */

package helper

import (
	"ETGo/components/base"
	"ETGo/conf"
	"github.com/jinzhu/gorm"
)

var (
	MysqlClient *gorm.DB
)

func InitMysql()  {
	dbConf := conf.Rconf.Mysql["test"]
	var err error
	MysqlClient, err = base.InitMysqlClient(dbConf)
	if err != nil {
		panic("mysql connect error: %v" + err.Error())
	}
}

func GetDb(tableName string) *gorm.DB {
	return MysqlClient.Table(tableName)
}