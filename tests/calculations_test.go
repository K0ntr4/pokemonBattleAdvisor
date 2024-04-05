package tests

import (
	"github.com/K0ntr4/pokemonBattleAdvisor/src"
	"slices"
	"testing"
)

func TestGetEnemyPokemonByName(t *testing.T) {
	testCases := []struct {
		name              string
		pokemonName       string
		expectedName      string
		expectedAbilities []string
		expectedTypes     []string
	}{
		{
			name:              "Test get enemy pokemon by name bulbasaur",
			pokemonName:       "bulbasaur",
			expectedName:      "bulbasaur",
			expectedAbilities: []string{"overgrow", "chlorophyll"},
			expectedTypes:     []string{"grass", "poison"},
		},
		{
			name:              "Test get enemy pokemon by name charmander",
			pokemonName:       "charmander",
			expectedName:      "charmander",
			expectedAbilities: []string{"blaze", "solar-power"},
			expectedTypes:     []string{"fire"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := pokemonbattleadvisor.GetEnemyPokemonByName(tc.pokemonName)
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}

			if actual.Name != tc.expectedName {
				t.Errorf("Expected name: %s, got: %s", tc.expectedName, actual.Name)
			}

			if len(actual.Abilities) != len(tc.expectedAbilities) {
				t.Errorf("Expected abilities length: %d, got: %d", len(tc.expectedAbilities), len(actual.Abilities))
			}

			for i, ability := range actual.Abilities {
				if !slices.Contains(tc.expectedAbilities, ability) {
					t.Errorf("Expected ability at index %d: %s, got: %s", i, tc.expectedAbilities[i], ability)
				}
			}

			if len(actual.Types) != len(tc.expectedTypes) {
				t.Errorf("Expected types length: %d, got: %d", len(tc.expectedTypes), len(actual.Types))
			}

			for i, typ := range actual.Types {
				if !slices.Contains(tc.expectedTypes, typ) {
					t.Errorf("Expected type at index %d: %s, got: %s", i, tc.expectedTypes[i], typ)
				}
			}
		})
	}
}

func TestBestMoveIndexAndEffectiveness(t *testing.T) {
	testCases := []struct {
		name          string
		pokemon       pokemonbattleadvisor.Pokemon
		enemy         pokemonbattleadvisor.Pokemon
		expectedIndex int
		expectedValue float64
	}{
		{
			name: "Test best move index and effectiveness",
			pokemon: pokemonbattleadvisor.Pokemon{
				Moves: []pokemonbattleadvisor.Move{
					{Name: "low-kick", Type: "fighting"},
					{Name: "false-swipe", Type: "normal"},
					{Name: "hail", Type: "ice"},
					{Name: "blizzard", Type: "ice"},
				},
			},
			enemy: pokemonbattleadvisor.Pokemon{
				Types: []string{"dark", "ice"},
			},
			expectedIndex: 0,
			expectedValue: 4.0,
		},
	}

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			actualIndex, actualValue := pokemonbattleadvisor.BestMoveIndexAndEffectiveness(&testCase.pokemon, &testCase.enemy)

			if actualIndex != testCase.expectedIndex {
				t.Errorf("Expected index: %d, got: %d", testCase.expectedIndex, actualIndex)
			}

			if actualValue != testCase.expectedValue {
				t.Errorf("Expected value: %f, got: %f", testCase.expectedValue, actualValue)
			}
		})
	}
}

func TestBestPokemonMoveAndShouldSwitch(t *testing.T) {
	testCases := []struct {
		name          string
		team          []pokemonbattleadvisor.Pokemon
		enemy         pokemonbattleadvisor.Pokemon
		expectedParty int
		expectedMove  int
		expectedValue bool
	}{
		{
			name: "Test best pokemon move and should switch",
			team: []pokemonbattleadvisor.Pokemon{
				{
					Moves: []pokemonbattleadvisor.Move{
						{Name: "poison-jab", Type: "poison"},
						{Name: "false-swipe", Type: "normal"},
						{Name: "hail", Type: "ice"},
						{Name: "blizzard", Type: "ice"},
					},
				},
				{
					Moves: []pokemonbattleadvisor.Move{
						{Name: "moonblast", Type: "fairy"},
						{Name: "flash", Type: "normal"},
						{Name: "flamethrower", Type: "fire"},
						{Name: "double-slap", Type: "normal"},
					},
				},
			},
			enemy: pokemonbattleadvisor.Pokemon{
				Types: []string{"dark", "ice"},
			},
			expectedParty: 1,
			expectedMove:  0,
			expectedValue: false,
		},
	}

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			actualParty, actualMove, actualValue := pokemonbattleadvisor.BestPokemonMoveAndShouldSwitch(&testCase.team, &testCase.enemy)

			if actualParty != testCase.expectedParty {
				t.Errorf("Expected party: %d, got: %d", testCase.expectedParty, actualParty)
			}

			if actualMove != testCase.expectedMove {
				t.Errorf("Expected move: %d, got: %d", testCase.expectedMove, actualMove)
			}

			if actualValue != testCase.expectedValue {
				t.Errorf("Expected value: %t, got: %t", testCase.expectedValue, actualValue)
			}
		})
	}
}
