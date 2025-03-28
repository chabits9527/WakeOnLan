package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("服务异常: %v\n", err)
				debug.PrintStack()
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 500,
					"msg":  "服务器内部错误",
					"data": err,
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
