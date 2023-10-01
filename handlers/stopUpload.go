package handlers

import (
	"fmt"
	"io"
	"jonnedu/hng_task5/services"
	"jonnedu/hng_task5/typ"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func StopStream(c *gin.Context) {
	videoID := c.Param("videoID")
	
	file, err := io.ReadAll(c.Request.Body)
	if err != nil{
		typ.ErrorResponse(c, 400, "Error write file format")
	}

	mergedData, err := services.AppendTwoWebm(VideoDataMap[videoID], file)
	if err != nil {
		typ.ErrorResponse(c, 500, "Unable to append the two byte arrays")
		return
	}
	VideoDataMap[videoID] = mergedData
	log.Println("Length of bytes",len(VideoDataMap[videoID]))

	fileName := fmt.Sprintf("%s.webm", videoID)
	filePath := filepath.Join("./uploads", fileName)
	err = services.SaveToDisk(VideoDataMap[videoID], videoID, fileName, filePath)
	if err != nil{
		typ.ErrorResponse(c, 500, fmt.Sprintf("Error could not save to disk:%s", err.Error()))
		return
	}
	// background service to get transcription
	go services.GetVidTranscription(fileName)

	baseUrl := os.Getenv("BASE_URL")
	viewVideoUrl := fmt.Sprintf("%s/video/%s.webm", baseUrl, videoID)

	data := map[string]interface{}{
		"video_id":  videoID,
		"video_url": viewVideoUrl,
	}
	typ.SuccessResponse(c, http.StatusAccepted, "File upload successful and stream ended.", data)
}


