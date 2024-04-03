package main

import (
	"fmt"
	"github.com/K0ntr4/pokemonBattleAdvisor/src"
)

func main() {
	var err error
	var party []pokemonbattleadvisor.Pokemon

	party, err = pokemonbattleadvisor.GetRandomParty(0, 493)
	if err != nil {
		panic(err)
	}
	pokemonbattleadvisor.PrintParty(&party)

	println("Enemy Pokemon:")
	var enemy pokemonbattleadvisor.Pokemon
	enemy, err = pokemonbattleadvisor.GetRandomEnemyPokemon(0, 493)
	if err != nil {
		panic(err)
	}
	pokemonbattleadvisor.PrintHelperStructsPokemon(&enemy)

	for _, move := range party[0].Moves {
		print(move.Name + " - " + move.Type + " - ")
		fmt.Printf("%f\n", move.EffectivenessAgainst(&enemy, &party[0].Abilities))
	}
}
