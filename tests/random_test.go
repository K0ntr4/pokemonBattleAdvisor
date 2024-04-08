package tests_test

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
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			enemy, err := pokemonbattleadvisor.PokemonByName(testCase.pokemonName)
			if err != nil {
				t.Errorf("Error: %v", err)
			}

			if enemy.Name != testCase.expectedName {
				t.Errorf("Expected name: %s, got: %s", testCase.expectedName, enemy.Name)
			}

			if len(enemy.Abilities) != len(testCase.expectedAbilities) {
				t.Errorf("Expected abilities length: %d, got: %d", len(testCase.expectedAbilities), len(enemy.Abilities))
			}

			for _, ability := range testCase.expectedAbilities {
				if !slices.Contains(enemy.Abilities, ability) {
					t.Errorf("Expected ability: %s, got: %s", ability, enemy.Abilities)
				}
			}

			if len(enemy.Types) != len(testCase.expectedTypes) {
				t.Errorf("Expected types length: %d, got: %d", len(testCase.expectedTypes), len(enemy.Types))
			}

			for _, typ := range testCase.expectedTypes {
				if !slices.Contains(enemy.Types, typ) {
					t.Errorf("Expected type: %s, got: %s", typ, enemy.Types)
				}
			}
		})
	}
}

func TestGetRandomEnemyPokemon(t *testing.T) {
	enemy, err := pokemonbattleadvisor.GetRandomEnemyPokemon(0, 493)
	if err != nil {
		t.Errorf("Error getting random enemy pokemon: %v", err)
	}

	if enemy.Name == "" {
		t.Errorf("Expected enemy name to not be empty")
	}

	if len(enemy.Abilities) == 0 {
		t.Errorf("Expected enemy abilities to not be empty")
	}

	if len(enemy.Moves) == 0 {
		t.Errorf("Expected enemy moves to not be empty")
	}

	if len(enemy.Types) == 0 {
		t.Errorf("Expected enemy types to not be empty")
	}

	for _, move := range enemy.Moves {
		if move.Name == "" {
			t.Errorf("Expected move name to not be empty")
		}

		if move.Type == "" {
			t.Errorf("Expected move type to not be empty")
		}
	}

	for _, typ := range enemy.Types {
		if typ == "" {
			t.Errorf("Expected type to not be empty")
		}
	}
}

func TestGetRandomTeam(t *testing.T) {
	team, err := pokemonbattleadvisor.GetRandomTeam([]int{0, 493})
	if err != nil {
		t.Errorf("Error getting random team: %v", err)
	}

	if len(*team) == 0 {
		t.Errorf("Expected team to not be empty")
	}

	for _, pokemon := range *team {
		if pokemon.Name == "" {
			t.Errorf("Expected pokemon name to not be empty")
		}

		if len(pokemon.Abilities) == 0 {
			t.Errorf("Expected pokemon abilities to not be empty")
		}

		if len(pokemon.Moves) == 0 {
			t.Errorf("Expected pokemon moves to not be empty")
		}

		if len(pokemon.Types) == 0 {
			t.Errorf("Expected pokemon types to not be empty")
		}

		for _, move := range pokemon.Moves {
			if move.Move.Name == "" {
				t.Errorf("Expected move name to not be empty")
			}
		}

		for _, typ := range pokemon.Types {
			if typ.Type.Name == "" {
				t.Errorf("Expected type to not be empty")
			}
		}
	}
}

func TestGetRandomParty(t *testing.T) {
	party, err := pokemonbattleadvisor.GetRandomParty(0, 493)
	if err != nil {
		t.Errorf("Error getting random party: %v", err)
	}

	if len(party) == 0 {
		t.Errorf("Expected party to not be empty")
	}

	for _, pokemon := range party {
		if pokemon.Name == "" {
			t.Errorf("Expected pokemon name to not be empty")
		}

		if len(pokemon.Abilities) == 0 {
			t.Errorf("Expected pokemon abilities to not be empty")
		}

		if len(pokemon.Moves) == 0 {
			t.Errorf("Expected pokemon moves to not be empty")
		}

		if len(pokemon.Types) == 0 {
			t.Errorf("Expected pokemon types to not be empty")
		}

		for _, move := range pokemon.Moves {
			if move.Name == "" {
				t.Errorf("Expected move name to not be empty")
			}

			if move.Type == "" {
				t.Errorf("Expected move type to not be empty")
			}
		}

		for _, typ := range pokemon.Types {
			if typ == "" {
				t.Errorf("Expected type to not be empty")
			}
		}
	}
}
