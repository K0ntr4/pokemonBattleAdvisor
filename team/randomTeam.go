package team

import (
	"math/rand"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetRandomTeam(bounds ...int) []structs.Pokemon {
	// Set default bounds if not provided
	switch len(bounds) {
	case 0:
		bounds = []int{0, 493} // Default to gen 4
	case 1:
		bounds = append(bounds, 493)
	}

	// Fetch list of Pokémon within the given bounds
	pokemonList, err := pokeapi.Resource("pokemon", bounds[0], bounds[1])
	if err != nil {
		return nil
	}

	var team []structs.Pokemon
	for i := 0; i < 6; i++ {
		randomIndex := rand.Intn(len(pokemonList.Results))
		pokemon, err := pokeapi.Pokemon(pokemonList.Results[randomIndex].Name)
		if err != nil {
			continue
		}
		team = append(team, pokemon)
	}

	return team
}
