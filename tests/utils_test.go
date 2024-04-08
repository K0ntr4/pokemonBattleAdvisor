package tests

import (
	pokemonbattleadvisor "github.com/K0ntr4/pokemonBattleAdvisor/src"
	"github.com/mtslzr/pokeapi-go/structs"
	"slices"
	"testing"
)

func TestGetPartyPokemon(t *testing.T) {
	testCases := []struct {
		name              string
		pokemonName       string
		abilities         []string
		moves             []string
		expectedName      string
		expectedAbilities []string
		expectedTypes     []string
		expectedMoves     []pokemonbattleadvisor.Move
	}{
		{
			name:              "Test get party pokemon weavile",
			pokemonName:       "weavile",
			abilities:         []string{"pressure"},
			moves:             []string{"poison-jab", "false-swipe", "hail", "blizzard"},
			expectedName:      "weavile",
			expectedAbilities: []string{"pressure"},
			expectedTypes:     []string{"dark", "ice"},
			expectedMoves: []pokemonbattleadvisor.Move{
				{Name: "poison-jab", Type: "poison", Damage: 80.0, Accuracy: 1.0},
				{Name: "false-swipe", Type: "normal", Damage: 40.0, Accuracy: 1.0},
				{Name: "hail", Type: "ice", Damage: 0.0, Accuracy: 0.0},
				{Name: "blizzard", Type: "ice", Damage: 110.0, Accuracy: 0.7},
			},
		},
		{
			name:              "Test get party pokemon clefable",
			pokemonName:       "clefable",
			abilities:         []string{"unaware"},
			moves:             []string{"moonblast", "flash", "flamethrower", "double-slap"},
			expectedName:      "clefable",
			expectedAbilities: []string{"unaware"},
			expectedTypes:     []string{"fairy"},
			expectedMoves: []pokemonbattleadvisor.Move{
				{Name: "moonblast", Type: "fairy", Damage: 95.0, Accuracy: 1.0},
				{Name: "flash", Type: "normal", Damage: 0.0, Accuracy: 1.0},
				{Name: "flamethrower", Type: "fire", Damage: 90.0, Accuracy: 1.0},
				{Name: "double-slap", Type: "normal", Damage: 45.0, Accuracy: 0.85},
			},
		},
		{
			name:              "Test get party pokemon azumarill",
			pokemonName:       "azumarill",
			abilities:         []string{"huge-power"},
			moves:             []string{"ice-beam", "play-rough", "surf", "hydro-pump"},
			expectedName:      "azumarill",
			expectedAbilities: []string{"huge-power"},
			expectedTypes:     []string{"water", "fairy"},
			expectedMoves: []pokemonbattleadvisor.Move{
				{Name: "ice-beam", Type: "ice", Damage: 90.0, Accuracy: 1.0},
				{Name: "play-rough", Type: "fairy", Damage: 90.0, Accuracy: 0.9},
				{Name: "surf", Type: "water", Damage: 90.0, Accuracy: 1.0},
				{Name: "hydro-pump", Type: "water", Damage: 110.0, Accuracy: 0.8},
			},
		},
		{
			name:              "Test get party pokemon charmander",
			pokemonName:       "charmander",
			abilities:         []string{"blaze"},
			moves:             []string{"ember", "flamethrower", "fire-blast", "fire-spin"},
			expectedName:      "charmander",
			expectedAbilities: []string{"blaze"},
			expectedTypes:     []string{"fire"},
			expectedMoves: []pokemonbattleadvisor.Move{
				{Name: "ember", Type: "fire", Damage: 40.0, Accuracy: 1.0},
				{Name: "flamethrower", Type: "fire", Damage: 90.0, Accuracy: 1.0},
				{Name: "fire-blast", Type: "fire", Damage: 110.0, Accuracy: 0.85},
				{Name: "fire-spin", Type: "fire", Damage: 35.0, Accuracy: 0.85},
			},
		},
	}

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			actual := pokemonbattleadvisor.GetPartyPokemon(testCase.pokemonName, testCase.abilities, testCase.moves)

			if actual.Name != testCase.expectedName {
				t.Errorf("Expected name: %s, got: %s", testCase.expectedName, actual.Name)
			}

			if len(actual.Abilities) != len(testCase.expectedAbilities) {
				t.Errorf("Expected abilities length: %d, got: %d", len(testCase.expectedAbilities), len(actual.Abilities))
			}

			for i, ability := range actual.Abilities {
				if ability != testCase.expectedAbilities[i] {
					t.Errorf("Expected ability at index %d: %s, got: %s", i, testCase.expectedAbilities[i], ability)
				}
			}

			if len(actual.Types) != len(testCase.expectedTypes) {
				t.Errorf("Expected types length: %d, got: %d", len(testCase.expectedTypes), len(actual.Types))
			}

			if len(actual.Moves) != len(testCase.expectedMoves) {
				t.Errorf("Expected moves length: %d, got: %d", len(testCase.expectedMoves), len(actual.Moves))
			}

			for _, move := range testCase.expectedMoves {
				if !slices.Contains(actual.Moves, move) {
					t.Errorf("Expected move: %v, got: %v", move, actual.Moves)
				}
			}
		})
	}
}

func TestGetHelperStructsMove(t *testing.T) {
	testCases := []struct {
		name             string
		moveName         string
		expectedName     string
		expectedType     string
		expectedDamage   float64
		expectedAccuracy float64
		expectedError    bool
	}{
		{
			name:             "Test get move by name tackle",
			moveName:         "tackle",
			expectedName:     "tackle",
			expectedType:     "normal",
			expectedDamage:   40.0,
			expectedAccuracy: 1.0,
			expectedError:    false,
		},
		{
			name:             "Test get move by name razor-leaf",
			moveName:         "razor-leaf",
			expectedName:     "razor-leaf",
			expectedType:     "grass",
			expectedDamage:   55.0,
			expectedAccuracy: 0.95,
			expectedError:    false,
		},
		{
			name:             "Test get move by name double-slap",
			moveName:         "double-slap",
			expectedName:     "double-slap",
			expectedType:     "normal",
			expectedDamage:   45.0,
			expectedAccuracy: 0.85,
			expectedError:    false,
		},
		{
			name:             "Test get move by name invalid",
			moveName:         "invalid",
			expectedName:     "",
			expectedType:     "",
			expectedDamage:   0.0,
			expectedAccuracy: 0.0,
			expectedError:    true,
		},
	}

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := pokemonbattleadvisor.GetHelperStructsMove(testCase.moveName)
			if err != nil {
				if !testCase.expectedError {
					t.Errorf("Expected no error, got %v", err)
				}
				return
			}

			if actual.Name != testCase.expectedName {
				t.Errorf("Expected name: %s, got: %s", testCase.expectedName, actual.Name)
			}

			if actual.Type != testCase.expectedType {
				t.Errorf("Expected type: %s, got: %s", testCase.expectedType, actual.Type)
			}

			if actual.Damage != testCase.expectedDamage {
				t.Errorf("Expected damage: %f, got: %f", testCase.expectedDamage, actual.Damage)
			}

			if actual.Accuracy != testCase.expectedAccuracy {
				t.Errorf("Expected accuracy: %f, got: %f", testCase.expectedAccuracy, actual.Accuracy)
			}
		})
	}
}

func TestGetHelperStructsTypes(t *testing.T) {
	testCases := []struct {
		name          string
		pokemonName   string
		expectedTypes []string
		expectedError bool
	}{
		{
			name:          "Test get types by name bulbasaur",
			pokemonName:   "bulbasaur",
			expectedTypes: []string{"grass", "poison"},
			expectedError: false,
		},
		{
			name:          "Test get types by name charmander",
			pokemonName:   "charmander",
			expectedTypes: []string{"fire"},
			expectedError: false,
		},
		{
			name:          "Test get types by name invalid",
			pokemonName:   "invalid",
			expectedTypes: []string{},
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := pokemonbattleadvisor.GetHelperStructsTypes(testCase.pokemonName)
			if err != nil {
				if !testCase.expectedError {
					t.Errorf("Expected no error, got %v", err)
				}
				return
			}

			if len(actual) != len(testCase.expectedTypes) {
				t.Errorf("Expected types length: %d, got: %d", len(testCase.expectedTypes), len(actual))
			}

			for i, typ := range actual {
				if typ != testCase.expectedTypes[i] {
					t.Errorf("Expected type at index %d: %s, got: %s", i, testCase.expectedTypes[i], typ)
				}
			}
		})
	}
}

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

			for i, ability := range tc.expectedAbilities {
				if !slices.Contains(actual.Abilities, ability) {
					t.Errorf("Expected ability at index %d: %s, got: %s", i, ability, actual.Abilities[i])
				}
			}

			if len(actual.Types) != len(tc.expectedTypes) {
				t.Errorf("Expected types length: %d, got: %d", len(tc.expectedTypes), len(actual.Types))
			}

			for i, typ := range tc.expectedTypes {
				if !slices.Contains(actual.Types, typ) {
					t.Errorf("Expected type at index %d: %s, got: %s", i, typ, actual.Types[i])
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
					URL:  "https://pokeapi.co/api/v2/move/33/",
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
			{
				Move: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "double-slap",
					URL:  "https://pokeapi.co/api/v2/move/3/",
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
			{Name: "tackle", Type: "normal", Damage: 40.0, Accuracy: 1.0},
			{Name: "razor-leaf", Type: "grass", Damage: 55.0, Accuracy: 0.95},
			{Name: "double-slap", Type: "normal", Damage: 45.0, Accuracy: 0.85},
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

	for _, move := range expected.Moves {
		if !slices.Contains(actual.Moves, move) {
			t.Errorf("Expected move: %v, got: %v", move, actual.Moves)
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
