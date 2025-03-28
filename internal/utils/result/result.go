package result

import "github.com/gin-gonic/gin"

func Success(data any) gin.H {
	return gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	}
}

func DefaultSuccess() gin.H {
	return gin.H{
		"code": 200,
		"msg":  "success",
	}
}
