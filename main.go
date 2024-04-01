package main

import (
	"github.com/K0ntr4/pokemon_battle_advisor/team"
)

func main() {
	team := team.GetRandomTeam(0, 493)
	for _, pokemon := range team {
		println(pokemon.Name)
	}
}
