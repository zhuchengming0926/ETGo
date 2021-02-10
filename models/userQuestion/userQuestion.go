

/**
 * @File: userQuestion.go
 * @Author: zhuchengming
 * @Description:
 * @Date: 2021/2/8 11:16
 */

package userQuestion

import (
	"ETGo/helper"
	"bytes"
	"fmt"
	"log"
	"strings"
	"time"
)

type UserQuestion struct {
	Id               uint64
	Uid              uint64
	Phone            string
	Uname            string
	IsDeleted        int
	Client           uint
	QuestionContent  string
	ScreenshotUrls   string
	Status           int
	Device           string
	OpUid            uint64
	OpName           string
	AppVersion       string
	CreateTime       int64
	UpdateTime       int64
}

const TblName = "tblUserQuestion"

func tableName() string {
	return  TblName
}

func AddRecord(param *UserQuestion) (uint64, error) {
	db := helper.GetDb(tableName())
	param.UpdateTime = time.Now().Unix()
	param.CreateTime = time.Now().Unix()
	err := db.Create(param).Error
	return param.Id, err
}

func DelRecord(id, opUid uint64, opName string) error {
	return helper.GetDb(tableName()).Where("id = ?", id).Updates(
		map[string]interface{}{
			"is_deleted":1,
			"update_time":time.Now().Unix(),
			"op_uid": opUid,
			"op_name": opName,
		}).Error
}

func BatchDelRecords(ids []uint64, opUid uint64, opName string) error {
	err := helper.GetDb(TblName).Where("id in (?)", ids).
		Updates(map[string]interface{}{
			"is_deleted":1,
			"update_time":time.Now().Unix(),
			"op_uid": opUid,
			"op_name": opName,
	}).Error
	return err
}

func BatchAddRecords(params []*UserQuestion) error {
	batchSqlStr, err := getBatchSqlStr(params)
	if err != nil {
		return err
	}
	return helper.GetDb(tableName()).Exec(batchSqlStr).Error

	//另一种方式
	//var result []UserQuestion
	//err = helper.GetDb(tableName()).Raw(batchSqlStr).Scan(&result).Error
	//fmt.Println(result)
	//return err
}

func getBatchSqlStr(insertDatas []*UserQuestion) (string, error) {
	cnt := len(insertDatas)
	if cnt == 0 {
		return "", nil
	}
	var buffer bytes.Buffer
	sql := "insert into `tblUserQuestion` (`uid`, `phone`, `uname`, `question_content`, `op_uid`, `op_name`, `create_time`, `update_time`) values"
	if _, err := buffer.WriteString(sql); err != nil {
		return "", err
	}
	for _, e := range insertDatas {
		buffer.WriteString(fmt.Sprintf("(%d, %s, %s, %s, %d, %s, %d, %d),", e.Uid, forMatStrField(e.Phone), forMatStrField(e.Uname),
				forMatStrField(e.QuestionContent), e.OpUid, forMatStrField(e.OpName), e.CreateTime, e.UpdateTime))
	}
	execSql := buffer.String()
	execSql = strings.TrimRight(execSql, ",") + ";"
	log.Printf("batch sql is %s", execSql)
	return execSql, nil
}

func forMatStrField(ori string) string {
	return fmt.Sprintf("'%s'", ori)
}

func UpdateRecord(id uint64, opUid uint64, opName string) error {
	err := helper.GetDb(TblName).Where("id = ?", id).Update("op_name", opName, "op_uid", opUid).Error
	return err
}

func GetRecordsByIds(ids []uint64) ([]UserQuestion, error) {
	var result []UserQuestion
	err := helper.GetDb(tableName()).Where("id in (?)", ids).Find(&result).Error

	//另一种用法
	//err := helper.GetDb(TblName).Find(&result, "id in (?)", ids).Error
	return result, err
}