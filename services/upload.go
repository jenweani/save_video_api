package services

import (
	"bytes"
	"errors"
	"fmt"

	// "log"

	"os/exec"
)

func AppendTwoByteArray(mp4Data1, mp4Data2 []byte) ([]byte, error){
	// Check if both arrays are not empty
	if len(mp4Data1) == 0 || len(mp4Data2) == 0 {
		return nil, errors.New("one or both arrays are empty")
	}

	// Extract the header from the first MP4 file (usually 32 bytes)
	headerSize := 32
	if len(mp4Data1) < headerSize {
		return nil, errors.New("MP4 data array 1 is too short to contain a header")
	}
	mp4Header := mp4Data1[:headerSize]

	// Check if the headers match (you may want to implement more checks here)
	if !bytes.Equal(mp4Header, mp4Data2[:headerSize]) {
		return nil, errors.New("MP4 headers do not match or are not compatible")
	}

	// Concatenate the video and audio data from both MP4 files while preserving the header
	mergedMP4Data := append(mp4Header, mp4Data1[headerSize:]...)
	mergedMP4Data = append(mergedMP4Data, mp4Data2[headerSize:]...)

	return mergedMP4Data, nil
}

func MergeTwoVids(firstVid, secondVid string, outputFilePath string) ( error) {
	// Define the FFmpeg command to merge two videos.
	ffmpegCmd := exec.Command(
		"ffmpeg",
		"-i", firstVid,
		"-i", secondVid,
		"-filter_complex", "[0:v][0:a][1:v][1:a]concat=n=2:v=1:a=1[vout][aout]",
		"-map", "[vout]",
		"-map", "[aout]",
        "-strict", "experimental",
		outputFilePath,
	)

	// Run the FFmpeg command.
	err := ffmpegCmd.Run()
	if err != nil {
		fmt.Println("Error running FFmpeg:", err)
		return err
	}

	return nil
}