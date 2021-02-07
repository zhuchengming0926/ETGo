/**************************************************************************
 * Copyright (c) 2020 Zuoyebang Inc. All Rights Reserved
 **************************************************************************/

/**
 * @File: http.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/7 20:06
 */

package router

import (
	v1 "ETGo/controllers/http/index/v1"
	"github.com/gin-gonic/gin"
)

func Http(router *gin.Engine)  {
	v1Group := router.Group("/et/v1")
	{
		v1Group.POST("/index", v1.GetIndex)
	}
}

