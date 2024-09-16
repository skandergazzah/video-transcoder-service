package service

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/skandergazzah/transcode/model"
)

// Define available resolutions
var Resolutions = []string{
	"1080p",
	"720p",
	"480p",
	"360p",
	"240p",
	"144p",
}

func TranscodeService(inputFilePath, outputDir, fileNameWithoutExt string) (successfulResults, failedResolutions []model.TrancodeResult) {
	var (
		sem   = make(chan struct{}, 2) // Semaphore to limit concurrent transcoding jobs to 2
		wg    sync.WaitGroup
		mutex sync.Mutex
	)

	// Iterate through each resolution
	for _, resName := range Resolutions {
		sem <- struct{}{}
		wg.Add(1)

		go func(resName string) {
			defer func() {
				<-sem // Release semaphore slot
				wg.Done()
			}()

			// Define the output file
			outputFile := filepath.Join(outputDir, fmt.Sprintf("%s_%s.mp4", fileNameWithoutExt, resName))

			// Call the shell script
			cmd := exec.Command("/app/transcode.sh", inputFilePath, resName, outputFile)
			fmt.Println("Running command:", cmd.String())
			cmdOutput, err := cmd.CombinedOutput()
			if err != nil {
				mutex.Lock()
				failedResolutions = append(failedResolutions, model.TrancodeResult{
					Resolution: resName,
					Err:        fmt.Sprintf("Error transcoding %s: %v, output: %s\n", resName, err, string(cmdOutput)),
					Success:    false,
				})
				mutex.Unlock()
			} else {
				fmt.Printf("Successfully transcoded to %s\n", resName)
				mutex.Lock()
				successfulResults = append(successfulResults, model.TrancodeResult{
					Resolution: resName,
					OutputPath: outputFile,
					Success:    true,
				})
				mutex.Unlock()
			}

		}(resName)
	}

	wg.Wait()
	return successfulResults, failedResolutions
}
