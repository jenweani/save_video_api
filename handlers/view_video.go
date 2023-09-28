package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewVideoHandler(c *gin.Context) {
	file_name := c.Param("filename")

	c.File(fmt.Sprintf("uploads/%s", file_name))
}

func VideoPageHandler(c *gin.Context){
	file_name := c.Param("filename")
	url_path := fmt.Sprintf("/api/video/%s", file_name)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Url": url_path,
	})
}