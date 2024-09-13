package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kylelemons/godebug/pretty"
	"github.com/skandergazzah/transcode/service"
)

func Transcode(c *gin.Context) {
	// Get the uploaded file
	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileName := file.Filename
	inputFilePath := filepath.Join("/app/uploads", fileName)

	// Save the uploaded file
	err = c.SaveUploadedFile(file, inputFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to save file %s [error: %s]", fileName, err.Error())})
		return
	}

	// Generate output directory
	fileNameWithoutExt := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	outputDir := filepath.Join("/app/uploads", fmt.Sprintf("%s_output", fileNameWithoutExt))
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to create folder %s [error: %s]", outputDir, err.Error())})
		return
	}

	// Start the transcoding process in the background
	// which can take a long time depending on the video size and the number of resolutions,
	// This ensures that the client doesn't have to wait for the transcoding to complete.

	go func() {
		successedResolution, failedResolutions := service.TranscodeService(inputFilePath, outputDir, fileNameWithoutExt)
		fmt.Printf("Transcoding completed for %s.\n", fileNameWithoutExt)
		pretty.Print("Success: ", successedResolution)
		pretty.Print("Failed: ", failedResolutions)
	}()

	// Immediately respond with HTTP 200 and a message
	c.JSON(http.StatusOK, gin.H{
		"message":                 "Transcoding in progress",
		"resolutions_in_progress": service.Resolutions,
	})
}
