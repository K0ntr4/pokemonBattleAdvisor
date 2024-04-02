package tests_test

import (
	"github.com/K0ntr4/pokemon_battle_advisor/src"
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
		t.Run(tc.name, func(t *testing.T) {
			moveType := tc.moveType
			scrappy := false
			pokemon_battle_advisor.ApplySpecialAbilities(&tc.attackerAbilities, &moveType, &scrappy)

			if moveType != tc.expectedMoveType {
				t.Errorf("Expected move type to be %s, got %s", tc.expectedMoveType, moveType)
			}

			if scrappy != tc.expectedScrappy {
				t.Errorf("Expected scrappy to be %t, got %t", tc.expectedScrappy, scrappy)
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
		t.Run(tc.name, func(t *testing.T) {
			result := 1.0
			pokemon_battle_advisor.CalculateTypeEffectiveness(&tc.enemyTypes, &tc.moveType, tc.scrappy, &result)
			if result != tc.expectedResult {
				t.Errorf("Expected effectiveness to be %f, got %f", tc.expectedResult, result)
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
		t.Run(tc.name, func(t *testing.T) {
			result := 1.0
			pokemon_battle_advisor.ApplyEnemyPokemonAbilities(&tc.enemyAbilities, &tc.moveType, &result)
			if result != tc.expectedResult {
				t.Errorf("Expected effectiveness to be %f, got %f", tc.expectedResult, result)
			}
		})
	}
}

func TestMove_EffectivenessAgainst(t *testing.T) {
	testCases := []struct {
		name           string
		move           pokemon_battle_advisor.Move
		pokemon        pokemon_battle_advisor.Pokemon
		attackerAbil   []string
		expectedResult float64
	}{
		{
			name: "Normal move against Normal type",
			move: pokemon_battle_advisor.Move{Name: "tackle", Type: "normal"},
			pokemon: pokemon_battle_advisor.Pokemon{
				Abilities: []string{},
				Moves:     []pokemon_battle_advisor.Move{},
				Types:     []string{"normal"},
				Name:      "Snorlax",
			},
			attackerAbil:   []string{},
			expectedResult: 1.0,
		},
		{
			name: "Fire move against Grass type",
			move: pokemon_battle_advisor.Move{Name: "ember", Type: "fire"},
			pokemon: pokemon_battle_advisor.Pokemon{
				Abilities: []string{},
				Moves:     []pokemon_battle_advisor.Move{},
				Types:     []string{"grass"},
				Name:      "Bulbasaur",
			},
			attackerAbil:   []string{},
			expectedResult: 2.0,
		},
		{
			name: "Electric move against Water type with ability Volt Absorb",
			move: pokemon_battle_advisor.Move{Name: "thunderbolt", Type: "electric"},
			pokemon: pokemon_battle_advisor.Pokemon{
				Abilities: []string{"volt-absorb"},
				Moves:     []pokemon_battle_advisor.Move{},
				Types:     []string{"water"},
				Name:      "Lanturn",
			},
			attackerAbil:   []string{},
			expectedResult: 0.0,
		},
		{
			name: "Fighting move against Ghost type with ability Scrappy",
			move: pokemon_battle_advisor.Move{Name: "karate-chop", Type: "fighting"},
			pokemon: pokemon_battle_advisor.Pokemon{
				Abilities: []string{},
				Moves:     []pokemon_battle_advisor.Move{},
				Types:     []string{"ghost"},
				Name:      "Gengar",
			},
			attackerAbil:   []string{"scrappy"},
			expectedResult: 1.0,
		},
		{
			name: "Fairy move against Dragon type with ability Wonder Guard",
			move: pokemon_battle_advisor.Move{Name: "moonblast", Type: "fairy"},
			pokemon: pokemon_battle_advisor.Pokemon{
				Abilities: []string{"wonder-guard"},
				Moves:     []pokemon_battle_advisor.Move{},
				Types:     []string{"dragon"},
				Name:      "Dragonite",
			},
			attackerAbil:   []string{},
			expectedResult: 2.0,
		},
		{
			name: "Water move against Bug type with ability Wonder Guard",
			move: pokemon_battle_advisor.Move{Name: "surf", Type: "water"},
			pokemon: pokemon_battle_advisor.Pokemon{
				Abilities: []string{"wonder-guard"},
				Moves:     []pokemon_battle_advisor.Move{},
				Types:     []string{"bug"},
				Name:      "Shuckle",
			},
			attackerAbil:   []string{},
			expectedResult: 0.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.move.EffectivenessAgainst(&tc.pokemon, &tc.attackerAbil)
			if result != tc.expectedResult {
				t.Errorf("Expected effectiveness to be %f, got %f", tc.expectedResult, result)
			}
		})
	}
}