

/**
 * @File: userQuestion_test.go
 * @Author: zhuchengming
 * @Description:
 * @Date: 2021/2/8 12:08
 */

package dsUserQuestion

import (
	"ETGo/models/userQuestion"
	"ETGo/test/mysqltest"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestAddUserQuestion(t *testing.T) {
	mysqltest.InitMysql()
	uid := uint64(1111)
	uname := "lhy's zyb"
	phone := "111111002222"
	opName := "lyc"
	opUid := uint64(2222)

	id, err := AddUserQuestion(uid, opUid, uname, phone, opName)
	fmt.Println(id, err)
}

func TestBatchAddUserQuestions(t *testing.T) {
	mysqltest.InitMysql()
	var params []*userQuestion.UserQuestion
	params = []*userQuestion.UserQuestion{
		&userQuestion.UserQuestion{
			Uid:             125,
			Phone:           "112141",
			Uname:           "lhy'zyb",
			IsDeleted:       0,
			Client:          0,
			QuestionContent: "",
			ScreenshotUrls:  "",
			Status:          0,
			Device:          "",
			OpUid:           443,
			OpName:          "ttt",
		},
		&userQuestion.UserQuestion{
			Uid:             136,
			Phone:           "1151551",
			Uname:           "woma'zyb",
			IsDeleted:       0,
			Client:          0,
			QuestionContent: "",
			ScreenshotUrls:  "",
			Status:          0,
			Device:          "",
			OpUid:           1551,
			OpName:          "",
		},
	}
	err :=BatchAddUserQuestions(params)
	fmt.Println(err)
}

func TestDelSingleUserQuestion(t *testing.T) {
	mysqltest.InitMysql()
	id := uint64(6)
	err := DelSingleUserQuestion(id, 222,"nidie")
	fmt.Println(err)
}

func TestBatchDelUserQuestions(t *testing.T) {
	mysqltest.InitMysql()
	ids := []uint64{1,2,3}
	err := BatchDelUserQuestions(ids, 222,"nidie")
	fmt.Println(err)
}

func TestGetUserQuestions(t *testing.T) {
	mysqltest.InitMysql()
	ids := []uint64{12,13}
	retMap, _ := GetUserQuestions(ids)
	for _,val := range retMap {
		fmt.Println(val.Uname)
		//fmt.Println(jsoniter.MarshalToString(val))
	}
}