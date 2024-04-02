package main

import (
	"fmt"

	"github.com/K0ntr4/pokemon_battle_advisor/helperStructs"
	"github.com/K0ntr4/pokemon_battle_advisor/src"
)

func main() {
	var party []helperStructs.Pokemon
	var err error
	party, err = src.GetRandomParty(0, 493)
	if err != nil {
		panic(err)
	}
	src.PrintHelperStructsPokemon(&party[0])

	println("Enemy Pokemon:")
	var enemy helperStructs.Pokemon
	enemy, err = src.GetRandomEnemyPokemon(0, 493)
	if err != nil {
		panic(err)
	}
	src.PrintHelperStructsPokemon(&enemy)

	for _, move := range party[0].Moves {
		print(move.Name + " - " + move.Type + " - ")
		fmt.Printf("%f\n", move.EffectivenessAgainst(&enemy, &party[0].Abilities))
	}
}
