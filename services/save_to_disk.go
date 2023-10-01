package services

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func SaveToDisk(file []byte, videoID, filename, filePath string) error {

	if err := os.MkdirAll("./temp", os.ModePerm); err != nil {
        return errors.New("error : Unable to create uploads directory")
    }

	tempFilePath := filepath.Join("./temp", fmt.Sprintf("%s.yul", videoID))
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return err
	}
	tempFile.Write(file)
	defer tempFile.Close()
	

	cmd := exec.Command("ffmpeg",
		"-i", tempFilePath,
		"-f", "webm",
        "-c:v", "copy",
        "-c:a", "copy",
		filePath,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}
	
	defer os.Remove(tempFilePath)
	return nil
}