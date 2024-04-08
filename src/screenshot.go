package pokemonbattleadvisor

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
	"os"
)

const defaultScreenshotPath = "screenshot.png"

func save(img *image.RGBA, filePath string) (err error) {
	var file *os.File
	file, err = os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		return err
	}
	return nil
}

func TakeScreenshot(monitorIndex int, boundaries ...int) (string, error) {
	var img *image.RGBA
	var err error

	if !(len(boundaries) == 4) {
		fmt.Printf("No correct boundaries provided, taking full screenshot\n")
		img, err = screenshot.CaptureDisplay(monitorIndex)
	} else {
		img, err = screenshot.CaptureRect(image.Rect(boundaries[0], boundaries[1], boundaries[2], boundaries[3]))
	}
	if err != nil {
		return "", err
	}

	err = save(img, defaultScreenshotPath)
	if err != nil {
		return "", err
	}
	return defaultScreenshotPath, nil
}
