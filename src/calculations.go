package pokemonbattleadvisor

import (
	"sort"
)

type TeamScore struct {
	score        MoveScore
	pokemonIndex int
}

type MoveScore struct {
	Eff       float64
	Damage    float64
	Acc       float64
	MoveIndex int
}

func isMoveBetter(m1, m2 MoveScore) bool {
	return m1.Eff*m1.Damage*m1.Acc > m2.Eff*float64(m2.Damage)*m2.Acc
}

func rankMoves(results *[]MoveScore) {
	sort.Slice(*results, func(i, j int) bool {
		return isMoveBetter((*results)[i], (*results)[j])
	})
}

func RankPokemonMoves(pokemon, enemy *Pokemon) []MoveScore {
	var results []MoveScore
	for i := 0; i < len(pokemon.Moves); i++ {
		results = append(results, MoveScore{MoveIndex: i, Damage: pokemon.Moves[i].Damage, Eff: pokemon.Moves[i].EffectivenessAgainst(enemy, &pokemon.Abilities), Acc: pokemon.Moves[i].Accuracy})
	}
	rankMoves(&results)
	return results
}

func BestPokemonMoveAndShouldSwitch(team *[]Pokemon, enemy *Pokemon) (partyIndex, moveIndex int, shouldSwitch bool) {
	teamResults := make([]TeamScore, len(*team))
	for i := 0; i < len(*team); i++ {
		teamResults[i] = TeamScore{score: RankPokemonMoves(&(*team)[i], enemy)[0], pokemonIndex: i}
	}
	firstPokemonResults := RankPokemonMoves(&(*team)[0], enemy)
	sort.Slice(teamResults, func(i, j int) bool {
		return isMoveBetter(teamResults[i].score, teamResults[j].score)
	})
	if teamResults[0].score.Eff*float64(teamResults[0].score.Damage)*teamResults[0].score.Acc < firstPokemonResults[0].Eff*float64(firstPokemonResults[0].Damage)*firstPokemonResults[0].Acc*1.5 {
		return 0, firstPokemonResults[0].MoveIndex, false
	}
	return teamResults[0].pokemonIndex, teamResults[0].score.MoveIndex, true
}
