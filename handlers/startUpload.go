package handlers

import (
	"fmt"
	"jonnedu/hng_task5/typ"
	"net/http"
	"os"
	// "path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
// var VideoDataMap map[string]int
var VideoDataMap map[string][]byte

func StartStream(c *gin.Context) {
	videoID := fmt.Sprintf("vid_%s", uuid.New())
	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
        c.JSON(500, gin.H{"error": "Unable to create uploads directory"})
        return 
    }

	// VideoDataMap = map[string]int{
	// 	videoID: 0,
	// }

	// filePath := filepath.Join("./uploads", fmt.Sprintf("%s_ch%d.mp4", videoID, 0))
	// _, err := os.Create(filePath)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": "Unable to create video chunk"})
    //     return 
	// }

	data := map[string]interface{}{
		"video_id": videoID,
	}
	typ.SuccessResponse(c, http.StatusOK, "started stream", data)
}