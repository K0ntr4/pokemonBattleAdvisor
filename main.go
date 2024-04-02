package main

import (
	"fmt"

	"github.com/K0ntr4/pokemon_battle_advisor/src"
)

func main() {
	var err error
	var party []pokemon_battle_advisor.Pokemon

	party, err = pokemon_battle_advisor.GetRandomParty(0, 493)
	if err != nil {
		panic(err)
	}
	pokemon_battle_advisor.PrintParty(&party)

	println("Enemy Pokemon:")
	var enemy pokemon_battle_advisor.Pokemon
	enemy, err = pokemon_battle_advisor.GetRandomEnemyPokemon(0, 493)
	if err != nil {
		panic(err)
	}
	pokemon_battle_advisor.PrintHelperStructsPokemon(&enemy)

	for _, move := range party[0].Moves {
		print(move.Name + " - " + move.Type + " - ")
		fmt.Printf("%f\n", move.EffectivenessAgainst(&enemy, &party[0].Abilities))
	}
}
