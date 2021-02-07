/**************************************************************************
 * Copyright (c) 2020 Zuoyebang Inc. All Rights Reserved
 **************************************************************************/

/**
 * @File: index.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/7 20:09
 */

package v1

import (
	"ETGo/helper"
	"github.com/gin-gonic/gin"
)

func GetIndex(ctx *gin.Context)  {
	helper.Render.RenderJsonSucc(ctx, "你是我的爷")
	return
}