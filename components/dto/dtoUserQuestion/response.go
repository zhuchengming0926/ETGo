

/**
 * @File: response.go
 * @Author: zhuchengming
 * @Description:
 * @Date: 2021/2/8 16:39
 */

package dtoUserQuestion

//问题反馈详情
type DetailRes struct {
	Id              uint64   `json:"id"`
	FeedBackTime    string   `json:"feedBackTime"`
	Phone           string   `json:"phone"`
	Uid             uint64   `json:"uid"`
	Uname           string   `json:"uname"`
	Content         string   `json:"content"`
	FeedBackImgUrls []string `json:"feedBackImgUrls"`
	Device          string   `json:"device"`
	AppVersion      string   `json:"appVersion"`
}