package pokemonbattleadvisor

import (
	"crypto/rand"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"math/big"
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

func SetDefaultBounds(bounds *[]int) []int {
	// Set default bounds if not provided
	switch len(*bounds) {
	case 0:
		*bounds = []int{0, 493} // Default to gen 4
	case 1:
		*bounds = append(*bounds, 493)
	}
	return *bounds
}

func GetRandomIndex[T any](list []T) int {
	result, err := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
	if err != nil {
		panic(err)
	}
	return int(result.Int64())
}

func PokemonByName(name string) (pokemon Pokemon, err error) {
	var res structs.Pokemon
	res, err = pokeapi.Pokemon(name)
	pokemon = CastToHelperStructsPokemon(&res)
	return pokemon, err
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
