package src

import (
	"math/rand"

	"github.com/K0ntr4/pokemon_battle_advisor/helperStructs"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func getRandomTeam(bounds []int) (team []structs.Pokemon, err error) {
	// Fetch list of Pokémon within the given bounds
	var pokemonList structs.Resource
	pokemonList, err = pokeapi.Resource("pokemon", bounds[0], bounds[1])
	if err != nil {
		return nil, err
	}

	var pokemon structs.Pokemon
	for i := 0; i < 6; i++ {
		var randomIndex int = rand.Intn(len(pokemonList.Results))
		pokemon, err = pokeapi.Pokemon(pokemonList.Results[randomIndex].Name)
		if err != nil {
			continue
		}
		team = append(team, pokemon)
	}
	return team, err
}

func GetRandomParty(bounds ...int) (party []helperStructs.Pokemon, err error) {
	var randomTeam []structs.Pokemon
	var randomIndex int

	// Set default bounds if not provided
	switch len(bounds) {
	case 0:
		bounds = []int{0, 493} // Default to gen 4
	case 1:
		bounds = append(bounds, 493)
	}

	randomTeam, err = getRandomTeam(bounds)
	for _, pokemon := range randomTeam {
		var p helperStructs.Pokemon
		p.Name = pokemon.Name
		randomIndex = rand.Intn(len(pokemon.Abilities))
		p.Ability = pokemon.Abilities[randomIndex].Ability.Name
		var i int
		for i = 0; i < 4 && i < len(pokemon.Moves); i++ {
			randomIndex = rand.Intn(len(pokemon.Moves))
			p.Moves = append(p.Moves, helperStructs.Move{Name: pokemon.Moves[randomIndex].Move.Name})
		}
		party = append(party, p)
	}
	return party, err
}

func PrintParty(party []helperStructs.Pokemon) {
	for _, pokemon := range party {
		println(pokemon.Name)
		for range pokemon.Name {
			print("-")
		}
		println()
		println("Ability: ")
		println(pokemon.Ability)
		println("Moves:")
		for _, move := range pokemon.Moves {
			println(move.Name)
		}
		println()
	}
}
