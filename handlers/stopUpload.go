package handlers

import (
	"fmt"
	"io"
	"jonnedu/hng_task5/services"
	"jonnedu/hng_task5/typ"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func StopStream(c *gin.Context) {
	videoID := c.Param("videoID")

	if err := os.MkdirAll("./temp", os.ModePerm); err != nil {
        c.JSON(500, gin.H{"error": "Unable to create uploads directory"})
        return 
    }
	
	file, err := io.ReadAll(c.Request.Body)
	if err != nil{
		typ.ErrorResponse(c, 400, "Error write file format")
	}

	mergedData, err := services.AppendTwoByteArray(VideoDataMap[videoID], file)
	if err != nil {
		typ.ErrorResponse(c, 500, "Unable to append the two byte arrays")
		return
	}
	VideoDataMap[videoID] = mergedData
	log.Println("Length of bytes",len(VideoDataMap[videoID]))

	fileName := fmt.Sprintf("%s.mp4", videoID)
	filePath := filepath.Join("./uploads", fileName)
	err = saveToDisk(VideoDataMap[videoID], videoID, fileName, filePath)
	if err != nil{
		typ.ErrorResponse(c, 500, fmt.Sprintf("Error could not save to disk:%s", err.Error()))
		return
	}

	baseUrl := os.Getenv("BASE_URL")
	viewVideoUrl := fmt.Sprintf("%s/video/%s.mp4", baseUrl, videoID)

	data := map[string]interface{}{
		"video_id":  videoID,
		"video_url": viewVideoUrl,
	}
	typ.SuccessResponse(c, http.StatusAccepted, "File upload successful and stream ended.", data)
}

func saveToDisk(file []byte, videoID, filename, filePath string) (error){

	tempFilePath := filepath.Join("./temp", fmt.Sprintf("%s.yul",videoID))
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return err
	}
	tempFile.Write(file)
	defer tempFile.Close()

	cmd := exec.Command("ffmpeg",
		"-i", tempFilePath,
		"-f", "s16le",
		"-ar", "44100",
		"-ac", "2",
		"-f", "mp4",
		"-vcodec", "libx264",
		"-preset", "medium",
		"-tune", "film",
		"-acodec", "aac",
		filePath,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
