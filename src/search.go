package pokemon_battle_advisor

import (
	"errors"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/sajari/fuzzy"
	"strings"
)

func normalizePokemonName(pokemonName *string) {
	*pokemonName = strings.ToLower(*pokemonName)
	*pokemonName = strings.ReplaceAll(*pokemonName, " ", "-")
	*pokemonName = strings.ReplaceAll(*pokemonName, ".", "")
	*pokemonName = strings.ReplaceAll(*pokemonName, "'", "")
	*pokemonName = strings.ReplaceAll(*pokemonName, ":", "")
	*pokemonName = strings.ReplaceAll(*pokemonName, "♀", "-f")
	*pokemonName = strings.ReplaceAll(*pokemonName, "♂", "-m")
	*pokemonName = strings.ReplaceAll(*pokemonName, "é", "e")
	*pokemonName = strings.ReplaceAll(*pokemonName, "%", "")
}

func FuzzySearchPokemon(pokemonName string, bounds ...int) (pokemon structs.Pokemon, err error) {
	SetDefaultBounds(&bounds)
	normalizePokemonName(&pokemonName)

	var pokemonList structs.Resource
	pokemonList, err = pokeapi.Resource("pokemon", bounds[0], bounds[1])
	if err != nil {
		return pokemon, err
	}

	var pokemonNameList []string = make([]string, len(pokemonList.Results))
	for i := 0; i < len(pokemonList.Results); i++ {
		pokemonNameList[i] = pokemonList.Results[i].Name
	}

	model := fuzzy.NewModel()
	model.SetThreshold(1)
	model.SetDepth(3)
	model.Train(pokemonNameList)

	var name string = model.SpellCheck(pokemonName)
	if name == "" {
		return structs.Pokemon{}, errors.New("no pokemon found")
	}
	return pokeapi.Pokemon(name)
}
