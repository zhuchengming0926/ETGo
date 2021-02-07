/**************************************************************************
 * Copyright (c) 2020 Zuoyebang Inc. All Rights Reserved
 **************************************************************************/

/**
 * @File: main.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/7 19:49
 */

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":8000")
}
