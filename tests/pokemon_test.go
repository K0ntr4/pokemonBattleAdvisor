package tests_test

import (
	"github.com/K0ntr4/pokemonBattleAdvisor/src"
	"testing"
)

func TestApplySpecialAbilities(t *testing.T) {
	testCases := []struct {
		name              string
		attackerAbilities []string
		moveType          string
		expectedMoveType  string
		expectedScrappy   bool
	}{
		{
			name:              "Normalize ability",
			attackerAbilities: []string{"normalize"},
			moveType:          "fire",
			expectedMoveType:  "normal",
			expectedScrappy:   false,
		},
		{
			name:              "Scrappy ability",
			attackerAbilities: []string{"scrappy"},
			moveType:          "normal",
			expectedMoveType:  "normal",
			expectedScrappy:   true,
		},
		{
			name:              "No abilities",
			attackerAbilities: []string{"overgrow", "chlorophyll"},
			moveType:          "water",
			expectedMoveType:  "water",
			expectedScrappy:   false,
		},
	}

	for _, tc := range testCases {
		testCase := tc // capture range variable
		t.Run(testCase.name, func(t *testing.T) {
			moveType := testCase.moveType
			scrappy := false
			pokemonbattleadvisor.ApplySpecialAbilities(&testCase.attackerAbilities, &moveType, &scrappy)

			if moveType != testCase.expectedMoveType {
				t.Errorf("Expected move type to be %s, got %s", testCase.expectedMoveType, moveType)
			}

			if scrappy != testCase.expectedScrappy {
				t.Errorf("Expected scrappy to be %t, got %t", testCase.expectedScrappy, scrappy)
			}
		})
	}
}

func TestCalculateTypeEffectiveness(t *testing.T) {
	testCases := []struct {
		name           string
		enemyTypes     []string
		moveType       string
		scrappy        bool
		expectedResult float64
	}{
		{
			name:           "Fire move against Water type",
			enemyTypes:     []string{"water"},
			moveType:       "fire",
			expectedResult: 0.5,
		},
		{
			name:           "Electric move against Electric type",
			enemyTypes:     []string{"electric"},
			moveType:       "electric",
			expectedResult: 0.5,
		},
		{
			name:           "Normal move against Ghost type",
			enemyTypes:     []string{"ghost"},
			moveType:       "normal",
			expectedResult: 0.0,
		},
		{
			name:           "Normal move against Ghost type with ability Scrappy",
			enemyTypes:     []string{"ghost"},
			moveType:       "normal",
			scrappy:        true,
			expectedResult: 1.0,
		},
	}

	for _, tc := range testCases {
		testCase := tc // capture range variable
		t.Run(testCase.name, func(t *testing.T) {
			result := 1.0
			pokemonbattleadvisor.CalculateTypeEffectiveness(&testCase.enemyTypes, &testCase.moveType, testCase.scrappy, &result)
			if result != testCase.expectedResult {
				t.Errorf("Expected effectiveness to be %f, got %f", testCase.expectedResult, result)
			}
		})
	}
}

func TestApplyEnemyPokemonAbilities(t *testing.T) {
	testCases := []struct {
		name           string
		enemyAbilities []string
		moveType       string
		expectedResult float64
	}{
		{
			name:           "Water type with ability Water Absorb",
			enemyAbilities: []string{"water-absorb"},
			moveType:       "water",
			expectedResult: 0.0,
		},
		{
			name:           "Fire type with ability Flash Fire",
			enemyAbilities: []string{"flash-fire"},
			moveType:       "fire",
			expectedResult: 0.0,
		},
		{
			name:           "Fire type with ability Heatproof",
			enemyAbilities: []string{"heatproof"},
			moveType:       "fire",
			expectedResult: 0.5,
		},
	}

	for _, tc := range testCases {
		testCase := tc // capture range variable
		t.Run(testCase.name, func(t *testing.T) {
			result := 1.0
			pokemonbattleadvisor.ApplyEnemyPokemonAbilities(&testCase.enemyAbilities, &testCase.moveType, &result)
			if result != testCase.expectedResult {
				t.Errorf("Expected effectiveness to be %f, got %f", testCase.expectedResult, result)
			}
		})
	}
}

func TestMove_EffectivenessAgainst(t *testing.T) {
	testCases := []struct {
		name           string
		move           pokemonbattleadvisor.Move
		pokemon        pokemonbattleadvisor.Pokemon
		attackerAbil   []string
		expectedResult float64
	}{
		{
			name: "Normal move against Normal type",
			move: pokemonbattleadvisor.Move{Name: "tackle", Type: "normal"},
			pokemon: pokemonbattleadvisor.Pokemon{
				Abilities: []string{},
				Moves:     []pokemonbattleadvisor.Move{},
				Types:     []string{"normal"},
				Name:      "Snorlax",
			},
			attackerAbil:   []string{},
			expectedResult: 1.0,
		},
		{
			name: "Fire move against Grass type",
			move: pokemonbattleadvisor.Move{Name: "ember", Type: "fire"},
			pokemon: pokemonbattleadvisor.Pokemon{
				Abilities: []string{},
				Moves:     []pokemonbattleadvisor.Move{},
				Types:     []string{"grass"},
				Name:      "Bulbasaur",
			},
			attackerAbil:   []string{},
			expectedResult: 2.0,
		},
		{
			name: "Electric move against Water type with ability Volt Absorb",
			move: pokemonbattleadvisor.Move{Name: "thunderbolt", Type: "electric"},
			pokemon: pokemonbattleadvisor.Pokemon{
				Abilities: []string{"volt-absorb"},
				Moves:     []pokemonbattleadvisor.Move{},
				Types:     []string{"water"},
				Name:      "Lanturn",
			},
			attackerAbil:   []string{},
			expectedResult: 0.0,
		},
		{
			name: "Fighting move against Ghost type with ability Scrappy",
			move: pokemonbattleadvisor.Move{Name: "karate-chop", Type: "fighting"},
			pokemon: pokemonbattleadvisor.Pokemon{
				Abilities: []string{},
				Moves:     []pokemonbattleadvisor.Move{},
				Types:     []string{"ghost"},
				Name:      "Gengar",
			},
			attackerAbil:   []string{"scrappy"},
			expectedResult: 1.0,
		},
		{
			name: "Fairy move against Dragon type with ability Wonder Guard",
			move: pokemonbattleadvisor.Move{Name: "moonblast", Type: "fairy"},
			pokemon: pokemonbattleadvisor.Pokemon{
				Abilities: []string{"wonder-guard"},
				Moves:     []pokemonbattleadvisor.Move{},
				Types:     []string{"dragon"},
				Name:      "Dragonite",
			},
			attackerAbil:   []string{},
			expectedResult: 2.0,
		},
		{
			name: "Water move against Bug type with ability Wonder Guard",
			move: pokemonbattleadvisor.Move{Name: "surf", Type: "water"},
			pokemon: pokemonbattleadvisor.Pokemon{
				Abilities: []string{"wonder-guard"},
				Moves:     []pokemonbattleadvisor.Move{},
				Types:     []string{"bug"},
				Name:      "Shuckle",
			},
			attackerAbil:   []string{},
			expectedResult: 0.0,
		},
	}

	for _, tc := range testCases {
		testCase := tc // capture range variable
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.move.EffectivenessAgainst(&testCase.pokemon, &testCase.attackerAbil)
			if result != testCase.expectedResult {
				t.Errorf("Expected effectiveness to be %f, got %f", testCase.expectedResult, result)
			}
		})
	}
}
