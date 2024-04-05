package tests

import (
	"github.com/K0ntr4/pokemonBattleAdvisor/src"
	"testing"
)

func TestClassify(t *testing.T) {
	testCases := []struct {
		name      string
		imagePath string
		expected  string
	}{
		{
			name:      "Test classify squirtle",
			imagePath: "testfiles/testImageSquirtle.png",
			expected:  "squirtle",
		},
		{
			name:      "Test classify sandshrew",
			imagePath: "testfiles/testImageSandshrew.png",
			expected:  "sandshrew",
		},
	}

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			classifierResult, err := pokemonbattleadvisor.Classify(testCase.imagePath)
			if err != nil {
				t.Errorf("Error: %s", err)
			}

			if classifierResult != testCase.expected {
				t.Errorf("Expected %s, got %s", testCase.expected, classifierResult)
			}
		})
	}
}
