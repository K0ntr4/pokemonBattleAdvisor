//nolint:dupl // This file contains tests for multiple functions that are similar
package tests

import (
	"github.com/K0ntr4/pokemonBattleAdvisor/src"
	"testing"
)

func TestRankPokemonMoves(t *testing.T) {
	testCases := []struct {
		name          string
		pokemon       pokemonbattleadvisor.Pokemon
		enemy         pokemonbattleadvisor.Pokemon
		expectedIndex int
		expectedValue float64
	}{
		{
			name: "Super effective move but no damage, best move is false-swipe",
			pokemon: pokemonbattleadvisor.Pokemon{
				Moves: []pokemonbattleadvisor.Move{
					{Name: "low-kick", Type: "fighting", Damage: 0.0, Accuracy: 1.0},
					{Name: "false-swipe", Type: "normal", Damage: 40.0, Accuracy: 1.0},
					{Name: "hail", Type: "ice", Damage: 0.0, Accuracy: 1.0},
					{Name: "blizzard", Type: "ice", Damage: 110.0, Accuracy: 0.7},
				},
			},
			enemy: pokemonbattleadvisor.Pokemon{
				Types: []string{"dark", "ice"},
			},
			expectedIndex: 1,
			expectedValue: 1.0,
		},
		{
			name: "Super effective move with damage, best move is low-kick",
			pokemon: pokemonbattleadvisor.Pokemon{
				Moves: []pokemonbattleadvisor.Move{
					{Name: "low-kick", Type: "fighting", Damage: 80.0, Accuracy: 1.0},
					{Name: "false-swipe", Type: "normal", Damage: 40.0, Accuracy: 1.0},
					{Name: "hail", Type: "ice", Damage: 0.0, Accuracy: 1.0},
					{Name: "blizzard", Type: "ice", Damage: 110.0, Accuracy: 0.7},
				},
			},
			enemy: pokemonbattleadvisor.Pokemon{
				Types: []string{"dark", "ice"},
			},
			expectedIndex: 0,
			expectedValue: 4.0,
		},
		{
			name: "All moves are not very effective, best move is blizzard",
			pokemon: pokemonbattleadvisor.Pokemon{
				Moves: []pokemonbattleadvisor.Move{
					{Name: "low-kick", Type: "fighting", Damage: 75.0, Accuracy: 1.0},
					{Name: "false-swipe", Type: "normal", Damage: 40.0, Accuracy: 1.0},
					{Name: "hail", Type: "ice", Damage: 0.0, Accuracy: 1.0},
					{Name: "blizzard", Type: "ice", Damage: 110.0, Accuracy: 0.7},
				},
			},
			enemy: pokemonbattleadvisor.Pokemon{
				Types: []string{"electric"},
			},
			expectedIndex: 3,
			expectedValue: 1.0,
		},
	}

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			var result = pokemonbattleadvisor.RankPokemonMoves(&testCase.pokemon, &testCase.enemy)
			if result[0].MoveIndex != testCase.expectedIndex {
				t.Errorf("Expected index: %d, got: %d", testCase.expectedIndex, result[0].MoveIndex)
			}

			if result[0].Eff != testCase.expectedValue {
				t.Errorf("Expected value: %f, got: %f", testCase.expectedValue, result[0].Eff)
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
			name: "Best pokemon against enemy dark and ice is second party member with first move",
			team: []pokemonbattleadvisor.Pokemon{
				{
					Moves: []pokemonbattleadvisor.Move{
						{Name: "poison-jab", Type: "poison", Damage: 80.0, Accuracy: 1.00},
						{Name: "false-swipe", Type: "normal", Damage: 40.0, Accuracy: 1.00},
						{Name: "hail", Type: "ice", Damage: 0.0, Accuracy: 100.0},
						{Name: "blizzard", Type: "ice", Damage: 110.0, Accuracy: 0.70},
					},
				},
				{
					Moves: []pokemonbattleadvisor.Move{
						{Name: "moonblast", Type: "fairy", Damage: 95.0, Accuracy: 1.00},
						{Name: "flash", Type: "normal", Damage: 0.0, Accuracy: 1.0},
						{Name: "flamethrower", Type: "fire", Damage: 90.0, Accuracy: 1.00},
						{Name: "double-slap", Type: "normal", Damage: 15.0, Accuracy: 0.85},
					},
				},
			},
			enemy: pokemonbattleadvisor.Pokemon{
				Types: []string{"dark", "ice"},
			},
			expectedParty: 1,
			expectedMove:  0,
			expectedValue: true,
		},
		{
			name: "Best pokemon against enemy electric is first party member with first move",
			team: []pokemonbattleadvisor.Pokemon{
				{
					Moves: []pokemonbattleadvisor.Move{
						{Name: "poison-jab", Type: "poison", Damage: 80.0, Accuracy: 1.00},
						{Name: "false-swipe", Type: "normal", Damage: 40.0, Accuracy: 1.00},
						{Name: "hail", Type: "ice", Damage: 0.0, Accuracy: 100.0},
						{Name: "blizzard", Type: "ice", Damage: 110.0, Accuracy: 0.70},
					},
				},
				{
					Moves: []pokemonbattleadvisor.Move{
						{Name: "moonblast", Type: "fairy", Damage: 95.0, Accuracy: 1.00},
						{Name: "flash", Type: "normal", Damage: 0.0, Accuracy: 1.0},
						{Name: "flamethrower", Type: "fire", Damage: 90.0, Accuracy: 1.00},
						{Name: "double-slap", Type: "normal", Damage: 15.0, Accuracy: 0.85},
					},
				},
			},
			enemy: pokemonbattleadvisor.Pokemon{
				Types: []string{"electric"},
			},
			expectedParty: 0,
			expectedMove:  0,
			expectedValue: false,
		},
		{
			name: "Best pokemon against enemy electric is second party member with first move",
			team: []pokemonbattleadvisor.Pokemon{
				{
					Moves: []pokemonbattleadvisor.Move{
						{Name: "moonblast", Type: "fairy", Damage: 95.0, Accuracy: 1.00},
						{Name: "flash", Type: "normal", Damage: 0.0, Accuracy: 1.0},
						{Name: "flamethrower", Type: "fire", Damage: 90.0, Accuracy: 1.00},
						{Name: "double-slap", Type: "normal", Damage: 15.0, Accuracy: 0.85},
					},
				},
				{
					Moves: []pokemonbattleadvisor.Move{
						{Name: "brick-break", Type: "fighting", Damage: 75.0, Accuracy: 1.00},
						{Name: "false-swipe", Type: "normal", Damage: 40.0, Accuracy: 1.00},
						{Name: "hail", Type: "ice", Damage: 0.0, Accuracy: 100.0},
						{Name: "blizzard", Type: "ice", Damage: 110.0, Accuracy: 0.70},
					},
				},
			},
			enemy: pokemonbattleadvisor.Pokemon{
				Types: []string{"electric"},
			},
			expectedParty: 0,
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
