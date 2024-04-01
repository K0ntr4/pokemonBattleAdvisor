package main

import (
	"math/rand"

	"github.com/K0ntr4/pokemon_battle_advisor/helperStructs"
	"github.com/K0ntr4/pokemon_battle_advisor/src"
	"github.com/mtslzr/pokeapi-go/structs"
)

func main() {
	var randomTeam []structs.Pokemon
	var err error
	randomTeam, err = src.GetRandomTeam(0, 493)
	if err != nil {
		println(err.Error())
		return
	}
	var party []helperStructs.Pokemon
	var randomIndex int
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
	for _, pokemon := range party {
		println(pokemon.Name)
		println(pokemon.Ability)
		for _, move := range pokemon.Moves {
			println(move.Name)
		}
	}
}
