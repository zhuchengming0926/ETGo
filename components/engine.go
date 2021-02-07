/**************************************************************************
 * Copyright (c) 2020 Zuoyebang Inc. All Rights Reserved
 **************************************************************************/

/**
 * @File: engine.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/7 20:45
 */

package components

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var engine *gin.Engine
var once sync.Once
func GetEngin() *gin.Engine {
	once.Do(func() {
		engine = gin.New()
	})
	return engine
}


