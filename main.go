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

	enemy, err = pokemonbattleadvisor.GetEnemyPokemonByName(classifierResult)
	if err != nil {
		panic(err)
	}
	pokemonbattleadvisor.PrintHelperStructsPokemon(&enemy)
	return enemy
}

func main() {
	var enemy = getEnemyPokemon()
	var party = []pokemonbattleadvisor.Pokemon{
		{
			Name: "weavile", Types: []string{"dark", "ice"},
			Abilities: []string{"pressure"},
			Moves: []pokemonbattleadvisor.Move{
				{Name: "poison-jab", Type: "poison"},
				{Name: "false-swipe", Type: "normal"},
				{Name: "hail", Type: "ice"},
				{Name: "blizzard", Type: "ice"},
			},
		},
		{
			Name: "clefable", Types: []string{"fairy"},
			Abilities: []string{"unaware"},
			Moves: []pokemonbattleadvisor.Move{
				{Name: "moonblast", Type: "fairy"},
				{Name: "flash", Type: "normal"},
				{Name: "flamethrower", Type: "fire"},
				{Name: "double-slap", Type: "normal"},
			},
		},
		{
			Name: "azumarill", Types: []string{"water", "fairy"},
			Abilities: []string{"huge-power"},
			Moves: []pokemonbattleadvisor.Move{
				{Name: "ice-beam", Type: "water"},
				{Name: "play-rough", Type: "fairy"},
				{Name: "surf", Type: "water"},
				{Name: "hydro-pump", Type: "fighting"},
			},
		},
		{
			Name: "luxray", Types: []string{"electric"},
			Abilities: []string{"rivalry"},
			Moves: []pokemonbattleadvisor.Move{
				{Name: "thunderbolt", Type: "electric"},
				{Name: "crunch", Type: "dark"},
				{Name: "flash", Type: "normal"},
				{Name: "discharge", Type: "electric"},
			},
		},
		{
			Name: "ludicolo", Types: []string{"water", "grass"},
			Abilities: []string{"swift-swim"},
			Moves: []pokemonbattleadvisor.Move{
				{Name: "dive", Type: "water"},
				{Name: "surf", Type: "water"},
				{Name: "giga-drain", Type: "grass"},
				{Name: "energy-ball", Type: "grass"},
			},
		},
		{
			Name: "sceptile", Types: []string{"grass"},
			Abilities: []string{"overgrow"},
			Moves: []pokemonbattleadvisor.Move{
				{Name: "cut", Type: "normal"},
				{Name: "dig", Type: "ground"},
				{Name: "energy-ball", Type: "grass"},
				{Name: "giga-drain", Type: "grass"},
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
