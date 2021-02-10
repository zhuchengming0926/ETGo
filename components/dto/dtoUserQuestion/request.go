/**************************************************************************
 *
 **************************************************************************/

/**
 * @File: request.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/8 16:40
 */

package dtoUserQuestion


//英语反馈问题详情接口
type DetailReq struct {
	Id uint64 `json:"id" binding:"required"`
}


