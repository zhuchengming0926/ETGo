/**************************************************************************
 *
 **************************************************************************/

/**
 * @File: userQuestion.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/8 16:38
 */

package svUserQuestion

import (
	"ETGo/components/dto/dtoUserQuestion"
	"ETGo/data/dsUserQuestion"
	"strings"
	"time"
)

//问题反馈详情,返回result不会为nil
func UserFeedBackDetail(id uint64) (*dtoUserQuestion.DetailRes, error) {
	recordMap, err := dsUserQuestion.GetUserQuestions([]uint64{id})
	if err != nil {
		return nil, err
	}
	if _, ok := recordMap[id]; !ok {
		return nil, nil
	}
	result := &dtoUserQuestion.DetailRes{} //初始化，不为nil
	result.Uname = recordMap[id].Uname
	result.Uid = recordMap[id].Uid
	result.Phone = recordMap[id].Phone
	result.Id = recordMap[id].Id
	result.FeedBackImgUrls = strings.Split(recordMap[id].ScreenshotUrls, ",")
	result.Content = recordMap[id].QuestionContent
	result.Device = recordMap[id].Device
	result.AppVersion = recordMap[id].AppVersion
	if recordMap[id].CreateTime > 0 {
		result.FeedBackTime = time.Unix(recordMap[id].CreateTime, 0).Format("2006/01/02 15:04:05")
	}
	return result, nil
}

