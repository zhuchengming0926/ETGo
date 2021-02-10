

/**
 * @File: userQuestion.go
 * @Author: zhuchengming
 * @Description:
 * @Date: 2021/2/8 12:05
 */

package dsUserQuestion

import (
	"ETGo/models/userQuestion"
	"fmt"
	"time"
)

func AddUserQuestion(uid, opUid uint64, uname, phone, opName string) (uint64, error) {
	params := userQuestion.UserQuestion{
		Uid:             uid,
		Phone:           phone,
		Uname:           uname,
		OpUid:           opUid,
		OpName:          opName,
	}

	id, err := userQuestion.AddRecord(&params)
	return id, err
}

func BatchAddUserQuestions(params []*userQuestion.UserQuestion) error {
	var insertDatas []*userQuestion.UserQuestion
	for i, val := range params {
		temp := &userQuestion.UserQuestion{
			Uid:             val.Uid,
			Phone:           val.Phone,
			Uname:           val.Uname,
			IsDeleted:       0,
			Client:          1,
			QuestionContent: fmt.Sprintf("这app咋这么难用，q——%d", i),
			ScreenshotUrls:  "",
			Status:          0,
			Device:          "",
			OpUid:           val.OpUid,
			OpName:          val.OpName,
			AppVersion:      "",
			CreateTime:      time.Now().Unix(),
			UpdateTime:      time.Now().Unix(),
		}
		insertDatas = append(insertDatas, temp)
	}
	err := userQuestion.BatchAddRecords(insertDatas)
	return err
}

func DelSingleUserQuestion(id uint64, opUid uint64, opName string) error {
	return userQuestion.DelRecord(id, opUid, opName)
}

func BatchDelUserQuestions(ids []uint64, opUid uint64, opName string) error {
	return userQuestion.BatchDelRecords(ids, opUid, opName)
}

func GetUserQuestions(ids []uint64) (map[uint64]*userQuestion.UserQuestion, error) {
	records, err :=  userQuestion.GetRecordsByIds(ids)
	if err != nil {
		return nil, err
	}
	retMap := make(map[uint64]*userQuestion.UserQuestion)
	for _, val := range records {
		temp := val //不能直接使用&val, 取得都是最后一条记录
		retMap[val.Id] = &temp
	}
	return retMap, nil
}