/**************************************************************************
 *
 **************************************************************************/

/**
 * @File: init.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/4 21:07
 */

package mysqltest

import (
	"ETGo/helper"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type MysqlTestConf struct {
	Service         string
	DataBase        string
	Addr            string
	User            string
	Password        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifeTime time.Duration
	ConnTimeOut     time.Duration
	WriteTimeOut    time.Duration
	ReadTimeOut     time.Duration
	LogMode         bool
}

func InitMysql()  {
	mysqlTestConf := MysqlTestConf{
		DataBase:        "zhuchengming",
		Addr:            "127.0.0.1:3306",
		User:            "root",
		Password:        "123456",
		MaxIdleConns:    10,
		MaxOpenConns:    1000,
		ConnMaxLifeTime: 3600,
		ConnTimeOut:     1500 * time.Millisecond,
		WriteTimeOut:    1500 * time.Millisecond,
		ReadTimeOut:     1500 * time.Millisecond,
		LogMode:         false,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=%s&readTimeout=%s&writeTimeout=%s&parseTime=True&loc=Asia%%2FShanghai",
		mysqlTestConf.User,
		mysqlTestConf.Password,
		mysqlTestConf.Addr,
		mysqlTestConf.DataBase,
		mysqlTestConf.ConnTimeOut,
		mysqlTestConf.ReadTimeOut,
		mysqlTestConf.WriteTimeOut)

	fmt.Println("连接字符串为：",dsn)
	client, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	helper.MysqlClient = client
}
