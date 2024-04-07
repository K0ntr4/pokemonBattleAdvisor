package pokemonbattleadvisor

import (
	"crypto/rand"
	"fmt"
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
	"math/big"
)

func GetHelperStructsTypes(pokemonName string) (types []string, err error) {
	var p structs.Pokemon
	p, err = pokeapi.Pokemon(pokemonName)
	if err != nil {
		return types, err
	}
	for i := 0; i < len(p.Types); i++ {
		types = append(types, p.Types[i].Type.Name)
	}
	return types, err
}

func GetHelperStructsMove(name string) (move Move, err error) {
	var m structs.Move
	m, err = pokeapi.Move(name)
	if err != nil {
		return move, err
	}
	var dmg float64
	if m.Meta.MaxHits != nil {
		minHits, ok := m.Meta.MinHits.(float64)
		if !ok {
			minHits = 1
		}
		maxHits, ok := m.Meta.MaxHits.(float64)
		if !ok {
			maxHits = 1
		}
		dmg = float64(m.Power) * float64((int(minHits)+int(maxHits))/2)
	} else {
		dmg = float64(m.Power)
	}
	move = Move{Name: m.Name, Type: m.Type.Name, Damage: dmg, Accuracy: float64(m.Accuracy) / 100}
	return move, err
}

func castToHelperStructsMove(pokemon *structs.Pokemon, castTo *Pokemon) {
	var move Move
	var err error
	var nMoves = len(pokemon.Moves)

	for i := 0; i < nMoves; i++ {
		var randomIndex = GetRandomIndex(pokemon.Moves)
		move, err = GetHelperStructsMove(pokemon.Moves[randomIndex].Move.Name)
		if err != nil {
			continue
		}
		castTo.Moves = append(castTo.Moves, move)
		pokemon.Moves = append(pokemon.Moves[:randomIndex], pokemon.Moves[randomIndex+1:]...)
	}
}

func CastToHelperStructsPokemon(pokemon *structs.Pokemon) (p Pokemon) {
	p.Name = pokemon.Name
	for i := 0; i < len(pokemon.Abilities); i++ {
		p.Abilities = append(p.Abilities, pokemon.Abilities[i].Ability.Name)
	}
	castToHelperStructsMove(pokemon, &p)
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

func PrintHelperStructsMove(move *Move) {
	println(move.Name)
	for range move.Name {
		print("-")
	}
	println()
	println("Type: ", move.Type)
	fmt.Printf("Damage: %.1f\n", move.Damage)
	fmt.Printf("Accuracy: %.1f\n", move.Accuracy)
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
	for i := 0; i < len(pokemon.Moves); i++ {
		PrintHelperStructsMove(&pokemon.Moves[i])
	}
	println()
}

func PrintParty(party *[]Pokemon) {
	for i := 0; i < len(*party); i++ {
		PrintHelperStructsPokemon(&(*party)[i])
	}
}
