package handlers

import (
	"jonnedu/hng_task5/typ"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoRouteHandler(c *gin.Context){
	typ.ErrorResponse(c, http.StatusNotFound, "Route not Found")
}