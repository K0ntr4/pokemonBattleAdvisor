package pokemonbattleadvisor

import (
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func CastToHelperStructsPokemon(pokemon *structs.Pokemon) (p Pokemon) {
	var randomIndex int

	p.Name = pokemon.Name
	for i := 0; i < len(pokemon.Abilities); i++ {
		p.Abilities = append(p.Abilities, pokemon.Abilities[i].Ability.Name)
	}
	var i int
	for i = 0; i < 4 && i < len(pokemon.Moves); i++ {
		randomIndex = GetRandomIndex(pokemon.Moves)
		move, err := pokeapi.Move(pokemon.Moves[randomIndex].Move.Name)
		if err != nil {
			continue
		}
		p.Moves = append(p.Moves, Move{Name: pokemon.Moves[randomIndex].Move.Name, Type: move.Type.Name})
	}
	for i := 0; i < len(pokemon.Types); i++ {
		p.Types = append(p.Types, pokemon.Types[i].Type.Name)
	}
	return p
}

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

func PrintHelperStructsPokemon(pokemon *Pokemon) {
	println(pokemon.Name)
	for range pokemon.Name {
		print("-")
	}
	println()
	println("Types: ")
	for _, t := range pokemon.Types {
		println(t)
	}
	println("Abilities: ")
	for _, ability := range pokemon.Abilities {
		println(ability)
	}
	println("Moves:")
	for _, move := range pokemon.Moves {
		println(move.Name)
	}
	println()
}

func PrintParty(party *[]Pokemon) {
	for i := 0; i < len(*party); i++ {
		PrintHelperStructsPokemon(&(*party)[i])
	}
}
