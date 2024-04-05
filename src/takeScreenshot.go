package pokemonbattleadvisor

import (
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
	"os"
)

const defaultScreenshotPath = "screenshot.png"

var cropBoarderByResolutions = map[[2]int]image.Rectangle{
	{3840, 2160}: image.Rect(1900, 700, 2400, 1000),
	{2560, 1440}: image.Rect(1250, 450, 1600, 800),
	{1920, 1080}: image.Rect(950, 350, 1200, 500),
	{1600, 900}:  image.Rect(800, 300, 1000, 400),
}

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

func getBounds(rectangle image.Rectangle) image.Rectangle {
	return cropBoarderByResolutions[[2]int{rectangle.Dx(), rectangle.Dy()}]
}

func TakeScreenshot() (filename string, err error) {
	var img *image.RGBA

	bounds := getBounds(screenshot.GetDisplayBounds(0))

	img, err = screenshot.CaptureRect(bounds)
	if err != nil {
		return "", err
	}

	err = save(img, defaultScreenshotPath)
	if err != nil {
		return "", err
	}
	return defaultScreenshotPath, nil
}
