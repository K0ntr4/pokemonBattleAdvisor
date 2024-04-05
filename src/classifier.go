package pokemonbattleadvisor

import (
	"os"
	"strings"

	"github.com/Kardbord/hfapigo/v2"
)

const ImageClassificationModel = "imjeffhi/pokemon_classifier"

func init() {
	key := os.Getenv("HUGGING_FACE_TOKEN")
	if key != "" {
		hfapigo.SetAPIKey(key)
	}
}

func Classify(imagePath string) (string, error) {
	response, err := hfapigo.SendImageClassificationRequest(ImageClassificationModel, imagePath)
	if err != nil {
		return "", err
	}

	return strings.ToLower(response[0].Label), nil
}
