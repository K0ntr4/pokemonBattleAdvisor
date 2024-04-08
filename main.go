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

	screenshot, err = pokemonbattleadvisor.TakeScreenshot(0, 1250, 450, 1600, 800)
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

func main() {
	var enemy = getEnemyPokemon()
	var party = []pokemonbattleadvisor.Pokemon{
		pokemonbattleadvisor.GetPartyPokemon("weavile", []string{"pressure"}, []string{"poison-jab", "false-swipe", "hail", "blizzard"}),
		pokemonbattleadvisor.GetPartyPokemon("clefable", []string{"unaware"}, []string{"moonblast", "flash", "flamethrower", "double-slap"}),
		pokemonbattleadvisor.GetPartyPokemon("azumarill", []string{"huge-power"}, []string{"ice-beam", "play-rough", "surf", "hydro-pump"}),
		pokemonbattleadvisor.GetPartyPokemon("luxray", []string{"rivalry"}, []string{"thunderbolt", "crunch", "flash", "discharge"}),
		pokemonbattleadvisor.GetPartyPokemon("ludicolo", []string{"swift-swim"}, []string{"dive", "surf", "giga-drain", "energy-ball"}),
		pokemonbattleadvisor.GetPartyPokemon("sceptile", []string{"overgrow"}, []string{"cut", "dig", "energy-ball", "giga-drain"}),
	}
	pokemonbattleadvisor.PrintParty(&party)

	partyMember, move, shouldSwitch := pokemonbattleadvisor.BestPokemonMoveAndShouldSwitch(&party, &enemy)
	if shouldSwitch {
		fmt.Println("Should switch")
		fmt.Printf("Best party member: %s\n", party[partyMember].Name)
		fmt.Println("Best move:")
		pokemonbattleadvisor.PrintHelperStructsMove(&party[partyMember].Moves[move])
	} else {
		fmt.Println("Should not switch")
		fmt.Println("Best move:")
		pokemonbattleadvisor.PrintHelperStructsMove(&party[partyMember].Moves[move])
	}
}
