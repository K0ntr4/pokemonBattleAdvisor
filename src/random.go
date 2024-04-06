package pokemonbattleadvisor

import (
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetRandomEnemyPokemon(bounds ...int) (pokemon Pokemon, err error) {
	SetDefaultBounds(&bounds)

	// Fetch list of Pokémon
	var pokemonList structs.Resource
	pokemonList, err = pokeapi.Resource("pokemon", bounds[0], bounds[1])
	if err != nil {
		return pokemon, err
	}

	var res structs.Pokemon
	res, err = pokeapi.Pokemon(pokemonList.Results[GetRandomIndex(pokemonList.Results)].Name)
	pokemon = CastToHelperStructsPokemon(&res)
	return pokemon, err
}

func GetRandomTeam(bounds []int) (team *[]structs.Pokemon, err error) {
	// Fetch list of Pokémon within the given bounds
	var pokemonList structs.Resource
	pokemonList, err = pokeapi.Resource("pokemon", bounds[0], bounds[1])
	if err != nil {
		return nil, err
	}

	var pokemon structs.Pokemon
	team = new([]structs.Pokemon)
	for i := 0; i < 6; i++ {
		randomIndex := GetRandomIndex(pokemonList.Results)
		pokemon, err = pokeapi.Pokemon(pokemonList.Results[randomIndex].Name)
		if err != nil {
			continue
		}
		*team = append(*team, pokemon)
	}
	return team, err
}

func GetRandomParty(bounds ...int) (party []Pokemon, err error) {
	var randomTeam *[]structs.Pokemon

	SetDefaultBounds(&bounds)

	randomTeam, err = GetRandomTeam(bounds)
	for i := 0; i < len(*randomTeam); i++ {
		var p = CastToHelperStructsPokemon(&(*randomTeam)[i])
		party = append(party, p)
	}
	return party, err
}
