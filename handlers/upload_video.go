package handlers

import (
	// "fmt"
	"io"
	"log"

	// "jonnedu/hng_task5/services"
	"jonnedu/hng_task5/services"
	"jonnedu/hng_task5/typ"
	"net/http"

	// "os"

	"github.com/gin-gonic/gin"
)

func StreamUpload(c *gin.Context)  {

	videoID := c.Param("videoID")

	// Copy the binary data to the video file
	file, err := io.ReadAll(c.Request.Body)
	if err != nil {
		typ.ErrorResponse(c, http.StatusInternalServerError, "Error: could not read file")
		return
	}

	if len(VideoDataMap[videoID]) < 5 {
		VideoDataMap[videoID] = append(VideoDataMap[videoID], file...)
		c.JSON(http.StatusOK, gin.H{"message": "Video stream received"})
		return
	}

	mergedData, err := services.AppendTwoByteArray(VideoDataMap[videoID], file)
	if err != nil {
		typ.ErrorResponse(c, 500, "Unable to append the two byte arrays")
		return
	}
	VideoDataMap[videoID] = mergedData
	log.Println("Length of bytes",len(VideoDataMap[videoID]))

	// VideoDataMap[videoID] += 1
	// appendID := VideoDataMap[videoID]
	// filePath := fmt.Sprintf("uploads/%s_ch%d.mp4", videoID, appendID)
	// os.WriteFile(filePath, file, 0644)

	// if appendID >= 2 {
	//     appendID := VideoDataMap[videoID]
	//     finalFilePath := fmt.Sprintf("uploads/%s_ch%d.mp4", videoID, appendID + 1)
	// 	firstVidPath := fmt.Sprintf("uploads/%s_ch%d.mp4", videoID, appendID - 1)
	// 	firstVid, err := os.Open(firstVidPath)
	// 	if err != nil {
	// 		typ.ErrorResponse(c, 500, "could not open pervious chunk")
	// 		return 
	// 	}
	//     err = services.MergeTwoVids(firstVid.Name(), filePath, finalFilePath)
	// 	if err != nil {
	// 		typ.ErrorResponse(c, 500, "could not merge two vids")
	// 		return
	// 	}
	// 	VideoDataMap[videoID] += 1
	// 	os.Remove(firstVidPath)
	// 	os.Remove(filePath)
	// }
	

	c.JSON(http.StatusOK, gin.H{"message": "Video stream received"})
}