

/**
 * @File: engine.go
 * @Author: zhuchengming
 * @Description:
 * @Date: 2021/2/7 20:45
 */

package components

import (
	"ETGo/middleware"
	"github.com/gin-gonic/gin"
	"sync"
)

var engine *gin.Engine
var once sync.Once
func GetEngin() *gin.Engine {
	once.Do(func() {
		engine = gin.New()
		engine.Use(middleware.Logger())
		engine.Use(middleware.Recover)
	})
	return engine
}


