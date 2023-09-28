package handlers

import (
	"jonnedu/hng_task5/typ"

	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context){
	typ.SuccessResponse(c, 200, "success", nil)
}