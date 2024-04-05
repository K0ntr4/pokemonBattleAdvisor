package tests_test

import (
	"github.com/K0ntr4/pokemonBattleAdvisor/src"
	"slices"
	"testing"

	"github.com/mtslzr/pokeapi-go/structs"
)

func TestCastToHelperStructsPokemon(t *testing.T) {
	p := structs.Pokemon{
		Abilities: []struct {
			Ability struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"ability"`
			IsHidden bool `json:"is_hidden"`
			Slot     int  `json:"slot"`
		}{
			{
				Ability: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "overgrow",
					URL:  "https://pokeapi.co/api/v2/ability/65/",
				},
				IsHidden: false,
				Slot:     1,
			},
			{
				Ability: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "sap-sipper",
					URL:  "https://pokeapi.co/api/v2/ability/157/",
				},
				IsHidden: false,
				Slot:     2,
			},
		},
		BaseExperience: 64,
		Forms: []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			{
				Name: "bulbasaur",
				URL:  "https://pokeapi.co/api/v2/pokemon-form/1/",
			},
		},
		GameIndices: []struct {
			GameIndex int `json:"game_index"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		}{
			{
				GameIndex: 1,
				Version: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "red",
					URL:  "https://pokeapi.co/api/v2/version/1/",
				},
			},
		},
		Height:    7,
		ID:        1,
		IsDefault: true,
		Moves: []struct {
			Move struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move"`
			VersionGroupDetails []struct {
				LevelLearnedAt  int `json:"level_learned_at"`
				MoveLearnMethod struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"move_learn_method"`
				VersionGroup struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"version_group"`
			} `json:"version_group_details"`
		}{
			{
				Move: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "tackle",
					URL:  "https://pokeapi.co/api/v2/move/1/",
				},
			},
			{
				Move: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "razor-leaf",
					URL:  "https://pokeapi.co/api/v2/move/75/",
				},
			},
		},
		Name:  "bulbasaur",
		Order: 1,
		Species: struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			Name: "bulbasaur",
			URL:  "https://pokeapi.co/api/v2/pokemon-species/1/",
		},
		Sprites: struct {
			BackDefault      string      `json:"back_default"`
			BackFemale       interface{} `json:"back_female"`
			BackShiny        string      `json:"back_shiny"`
			BackShinyFemale  interface{} `json:"back_shiny_female"`
			FrontDefault     string      `json:"front_default"`
			FrontFemale      interface{} `json:"front_female"`
			FrontShiny       string      `json:"front_shiny"`
			FrontShinyFemale interface{} `json:"front_shiny_female"`
		}{
			FrontDefault: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png",
		},
		Stats: []struct {
			BaseStat int `json:"base_stat"`
			Effort   int `json:"effort"`
			Stat     struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"stat"`
		}{
			{
				BaseStat: 45,
				Effort:   0,
				Stat: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "",
					URL:  "",
				},
			},
		},
		Types: []struct {
			Slot int `json:"slot"`
			Type struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"type"`
		}{
			{
				Slot: 1,
				Type: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "grass",
					URL:  "https://pokeapi.co/api/v2/type/12/",
				},
			},
			{
				Slot: 2,
				Type: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "poison",
					URL:  "https://pokeapi.co/api/v2/type/4/",
				},
			},
		},
		Weight: 69,
	}

	expected := pokemonbattleadvisor.Pokemon{
		Name:      "bulbasaur",
		Abilities: []string{"overgrow", "sap-sipper"},
		Moves: []pokemonbattleadvisor.Move{
			{Name: "tackle", Type: "normal"},
			{Name: "razor-leaf", Type: "grass"},
		},
		Types: []string{"grass", "poison"},
	}

	actual := pokemonbattleadvisor.CastToHelperStructsPokemon(&p)

	if actual.Name != expected.Name {
		t.Errorf("Expected name: %s, got: %s", expected.Name, actual.Name)
	}

	if len(actual.Abilities) != len(expected.Abilities) {
		t.Errorf("Expected abilities length: %d, got: %d", len(expected.Abilities), len(actual.Abilities))
	}

	for i, ability := range actual.Abilities {
		if ability != expected.Abilities[i] {
			t.Errorf("Expected ability at index %d: %s, got: %s", i, expected.Abilities[i], ability)
		}
	}

	if len(actual.Moves) != len(expected.Moves) {
		t.Errorf("Expected moves length: %d, got: %d", len(expected.Moves), len(actual.Moves))
	}

	for _, move := range actual.Moves {
		if !slices.Contains(expected.Moves, move) {
			t.Errorf("Expected move: %v, got: %v", move, expected.Moves)
		}
	}

	if len(actual.Types) != len(expected.Types) {
		t.Errorf("Expected types length: %d, got: %d", len(expected.Types), len(actual.Types))
	}

	for i, typ := range actual.Types {
		if typ != expected.Types[i] {
			t.Errorf("Expected type at index %d: %s, got: %s", i, expected.Types[i], typ)
		}
	}
}

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
			enemy, err := pokemonbattleadvisor.GetEnemyPokemonByName(testCase.pokemonName)
			if err != nil {
				t.Errorf("Error: %v", err)
			}

			if enemy.Name != testCase.expectedName {
				t.Errorf("Expected name: %s, got: %s", testCase.expectedName, enemy.Name)
			}

			if len(enemy.Abilities) != len(testCase.expectedAbilities) {
				t.Errorf("Expected abilities length: %d, got: %d", len(testCase.expectedAbilities), len(enemy.Abilities))
			}

			for _, ability := range enemy.Abilities {
				if !slices.Contains(testCase.expectedAbilities, ability) {
					t.Errorf("Expected ability: %s, got: %s", testCase.expectedAbilities, ability)
				}
			}

			if len(enemy.Types) != len(testCase.expectedTypes) {
				t.Errorf("Expected types length: %d, got: %d", len(testCase.expectedTypes), len(enemy.Types))
			}

			for _, typ := range enemy.Types {
				if !slices.Contains(testCase.expectedTypes, typ) {
					t.Errorf("Expected type: %s, got: %s", testCase.expectedTypes, typ)
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
