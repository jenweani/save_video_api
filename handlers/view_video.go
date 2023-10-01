package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ViewVideoHandler(c *gin.Context) {
	videoID := c.Param("videoID")

	c.File(fmt.Sprintf("uploads/%s.webm", videoID))
}

func ViewTranscriptHandler(c *gin.Context){
	file_name := c.Param("filename")

	c.File(fmt.Sprintf("transcripts/%s.vtt", file_name))
}