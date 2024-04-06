package pokemonbattleadvisor

import (
	"sort"
)

type TeamResult struct {
	score        MoveResult
	pokemonIndex int
}

type MoveResult struct {
	Eff       float64
	Damage    int
	Acc       float64
	MoveIndex int
}

func isMoveBetter(m1, m2 MoveResult) bool {
	return m1.Eff*float64(m1.Damage)*m1.Acc > m2.Eff*float64(m2.Damage)*m2.Acc
}

func rankMoves(results *[]MoveResult) {
	sort.Slice(*results, func(i, j int) bool {
		return isMoveBetter((*results)[i], (*results)[j])
	})
}

func RankPokemonMoves(pokemon, enemy *Pokemon) []MoveResult {
	var results []MoveResult
	for i := 0; i < len(pokemon.Moves); i++ {
		results = append(results, MoveResult{MoveIndex: i, Damage: pokemon.Moves[i].Damage, Eff: pokemon.Moves[i].EffectivenessAgainst(enemy, &pokemon.Abilities), Acc: pokemon.Moves[i].Accuracy})
	}
	rankMoves(&results)
	return results
}

func BestPokemonMoveAndShouldSwitch(team *[]Pokemon, enemy *Pokemon) (partyIndex, moveIndex int, shouldSwitch bool) {
	teamResults := make([]TeamResult, len(*team))
	for i := 0; i < len(*team); i++ {
		teamResults[i] = TeamResult{score: RankPokemonMoves(&(*team)[i], enemy)[0], pokemonIndex: i}
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
