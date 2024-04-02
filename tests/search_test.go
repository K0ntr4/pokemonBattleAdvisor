package tests

import (
	"errors"
	pokemon_battle_advisor "github.com/K0ntr4/pokemon_battle_advisor/src"
	"testing"
)

func TestFuzzySearchPokemon(t *testing.T) {
	testCases := []struct {
		name         string
		actualName   string
		expectedName string
		expectedErr  error
	}{
		{
			name:         "Exact match",
			actualName:   "charizard",
			expectedName: "charizard",
			expectedErr:  nil,
		},
		{
			name:         "Close match",
			actualName:   "charzard",
			expectedName: "charizard",
			expectedErr:  nil,
		},
		{
			name:         "distant match",
			actualName:   "mr. m(im@",
			expectedName: "mr-mime",
			expectedErr:  nil,
		},
		{
			name:         "error",
			actualName:   ".-,.-2131231.,445,34,56435342",
			expectedName: "",
			expectedErr:  errors.New("no pokemon found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pokemon, err := pokemon_battle_advisor.FuzzySearchPokemon(tc.actualName)
			if tc.expectedErr == nil && err != nil {
				t.Errorf("Expected no error, got %s for %s", err, tc.expectedName)
			} else if tc.expectedErr != nil && err == nil {
				t.Errorf("Expected error %s, got nil for %s", tc.expectedErr, tc.expectedName)
			} else if tc.expectedErr != nil && err != nil && tc.expectedErr.Error() != err.Error() {
				t.Errorf("Expected error %s, got %s for %s", tc.expectedErr, err, tc.expectedName)
			}

			if pokemon.Name != tc.expectedName {
				t.Errorf("Expected pokemon name to be %s, got %s", tc.expectedName, pokemon.Name)
			}
		})
	}
}
