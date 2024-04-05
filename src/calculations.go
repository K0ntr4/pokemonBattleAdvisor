package pokemonbattleadvisor

import (
	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

func GetEnemyPokemonByName(name string) (pokemon Pokemon, err error) {
	var res structs.Pokemon
	res, err = pokeapi.Pokemon(name)
	pokemon = CastToHelperStructsPokemon(&res)
	return pokemon, err
}

func BestMoveIndexAndEffectiveness(pokemon, enemy *Pokemon) (index int, highest float64) {
	for i := 0; i < len(pokemon.Moves); i++ {
		res := pokemon.Moves[i].EffectivenessAgainst(enemy, &pokemon.Abilities)
		if res > highest {
			highest = res
			index = i
		}
	}
	return index, highest
}

func BestPokemonMoveAndShouldSwitch(team *[]Pokemon, enemy *Pokemon) (partyIndex, moveIndex int, shouldSwitch bool) {
	var current, highest float64
	var currentMoveIndex int

	for i := len(*team) - 1; i >= 0; i-- {
		currentMoveIndex, current = BestMoveIndexAndEffectiveness(&(*team)[i], enemy)
		if current > highest {
			highest = current
			partyIndex = i
			moveIndex = currentMoveIndex
		}
	}
	if current < highest*0.5 {
		shouldSwitch = true
	} else {
		shouldSwitch = false
		moveIndex = currentMoveIndex
	}
	return partyIndex, moveIndex, shouldSwitch
}
