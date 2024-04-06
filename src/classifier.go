package pokemonbattleadvisor

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Kardbord/hfapigo/v2"
)

const ImageClassificationModel = "imjeffhi/pokemon_classifier"

func init() {
	key := os.Getenv("HUGGING_FACE_TOKEN")
	if key != "" {
		hfapigo.SetAPIKey(key)
	}
}

func getEstimatedTime(str string) int64 {
	re := regexp.MustCompile(`estimated_time":(\d+)`)
	match := re.FindStringSubmatch(str)
	if len(match) >= 2 {
		if estimatedTime, err := strconv.ParseInt(match[1], 10, 64); err == nil {
			return estimatedTime + 1
		}
	}
	return 3
}

func Classify(imagePath string) (string, error) {
	for i := 0; i < 3; i++ {
		response, err := hfapigo.SendImageClassificationRequest(ImageClassificationModel, imagePath)
		if err != nil && strings.Contains(err.Error(), "currently loading") {
			time.Sleep(time.Duration(getEstimatedTime(err.Error())) * time.Second)
			continue
		}
		return strings.ToLower(response[0].Label), nil
	}
	return "", errors.New("failed to classify image")
}
