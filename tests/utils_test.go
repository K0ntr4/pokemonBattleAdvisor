package tests

import (
	pokemonbattleadvisor "github.com/K0ntr4/pokemonBattleAdvisor/src"
	"github.com/mtslzr/pokeapi-go/structs"
	"slices"
	"testing"
)

func TestPokemonByName(t *testing.T) {
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
			actual, err := pokemonbattleadvisor.PokemonByName(tc.pokemonName)
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