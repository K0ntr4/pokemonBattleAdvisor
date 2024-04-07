package main

import (
	"encoding/json"
	"fmt"
	"github.com/K0ntr4/pokemonBattleAdvisor/src"
	"time"
)

type currentlyLoadingError struct {
	Error         string  `json:"error"`
	EstimatedTime float64 `json:"estimated_time"`
}

func getEnemyPokemon() (enemy pokemonbattleadvisor.Pokemon) {
	var err error
	var screenshot string
	var classifierResult string

	screenshot, err = pokemonbattleadvisor.TakeScreenshot()
	if err != nil {
		return
	}

	for i := 0; i < 3; i++ {
		classifierResult, err = pokemonbattleadvisor.Classify(screenshot)
		if err != nil {
			unmarshalErr := currentlyLoadingError{}
			_ = json.Unmarshal([]byte(err.Error()), &unmarshalErr)
			fmt.Printf("Estimated time: %f\n", unmarshalErr.EstimatedTime)
			time.Sleep(3 * time.Second)
		}
	}

	enemy, err = pokemonbattleadvisor.PokemonByName(classifierResult)
	if err != nil {
		panic(err)
	}
	pokemonbattleadvisor.PrintHelperStructsPokemon(&enemy)
	return enemy
}

func getMoveIgnoreError(name string) pokemonbattleadvisor.Move {
	move, err := pokemonbattleadvisor.GetHelperStructsMove(name)
	if err != nil {
		return pokemonbattleadvisor.Move{}
	}
	return move
}

func getTypesIgnoreError(pokemonName string) []string {
	types, err := pokemonbattleadvisor.GetHelperStructsTypes(pokemonName)
	if err != nil {
		return []string{}
	}
	return types
}

func main() {
	var enemy = getEnemyPokemon()
	var party = []pokemonbattleadvisor.Pokemon{
		{
			Name: "weavile", Types: getTypesIgnoreError("weavile"),
			Abilities: []string{"pressure"},
			Moves: []pokemonbattleadvisor.Move{
				getMoveIgnoreError("poison-jab"),
				getMoveIgnoreError("false-swipe"),
				getMoveIgnoreError("hail"),
				getMoveIgnoreError("blizzard"),
			},
		},
		{
			Name: "clefable", Types: getTypesIgnoreError("clefable"),
			Abilities: []string{"unaware"},
			Moves: []pokemonbattleadvisor.Move{
				getMoveIgnoreError("moonblast"),
				getMoveIgnoreError("flash"),
				getMoveIgnoreError("flamethrower"),
				getMoveIgnoreError("double-slap"),
			},
		},
		{
			Name: "azumarill", Types: getTypesIgnoreError("azumarill"),
			Abilities: []string{"huge-power"},
			Moves: []pokemonbattleadvisor.Move{
				getMoveIgnoreError("ice-beam"),
				getMoveIgnoreError("play-rough"),
				getMoveIgnoreError("surf"),
				getMoveIgnoreError("hydro-pump"),
			},
		},
		{
			Name: "luxray", Types: getTypesIgnoreError("luxray"),
			Abilities: []string{"rivalry"},
			Moves: []pokemonbattleadvisor.Move{
				getMoveIgnoreError("thunderbolt"),
				getMoveIgnoreError("crunch"),
				getMoveIgnoreError("flash"),
				getMoveIgnoreError("discharge"),
			},
		},
		{
			Name: "ludicolo", Types: getTypesIgnoreError("ludicolo"),
			Abilities: []string{"swift-swim"},
			Moves: []pokemonbattleadvisor.Move{
				getMoveIgnoreError("dive"),
				getMoveIgnoreError("surf"),
				getMoveIgnoreError("giga-drain"),
				getMoveIgnoreError("energy-ball"),
			},
		},
		{
			Name: "sceptile", Types: getTypesIgnoreError("sceptile"),
			Abilities: []string{"overgrow"},
			Moves: []pokemonbattleadvisor.Move{
				getMoveIgnoreError("cut"),
				getMoveIgnoreError("dig"),
				getMoveIgnoreError("energy-ball"),
				getMoveIgnoreError("giga-drain"),
			},
		},
	}
	pokemonbattleadvisor.PrintParty(&party)

	partyMember, move, shouldSwitch := pokemonbattleadvisor.BestPokemonMoveAndShouldSwitch(&party, &enemy)
	if shouldSwitch {
		fmt.Println("Should switch")
		fmt.Printf("Best party member: %s\n", party[partyMember].Name)
		fmt.Printf("Best move: %s\n", party[partyMember].Moves[move].Name)
	} else {
		fmt.Println("Should not switch")
		fmt.Printf("Best move: %s\n", party[0].Moves[move].Name)
	}
}
