package typ

import "github.com/gin-gonic/gin"

func SuccessResponse(c *gin.Context, code int, message string, data map[string]interface{}){
	c.JSON(code, gin.H{
		"status": "success",
		"message": message,
		"data": data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string){
	c.JSON(code, gin.H{
		"status": "error",
		"message": message,
	})
}