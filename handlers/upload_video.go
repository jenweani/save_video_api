package handlers

import (
	"fmt"
	"jonnedu/hng_task5/typ"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadVideoHandler(c *gin.Context)  {
	base_url := os.Getenv("BASE_URL")

	file, err := c.FormFile("file")
	if err != nil {
		typ.ErrorResponse(c, http.StatusBadRequest, "Could not read file")
		return
	}

	fileName := fmt.Sprintf("uploads/%s", file.Filename)

	err = os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		typ.ErrorResponse(c, http.StatusInternalServerError, "Could not create required directory")
		return
	}

	if err := c.SaveUploadedFile(file, fileName); err != nil {
		typ.ErrorResponse(c, http.StatusInternalServerError, "Could not save file to disk")
		return
	}

	viewFileUrl := fmt.Sprintf("%s/api/video/%s", base_url, file.Filename)

	typ.SuccessResponse(c, http.StatusOK, "File upload successful", map[string]interface{}{"video_loc_url": viewFileUrl})
}