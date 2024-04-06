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

func main() {
	var enemy = getEnemyPokemon()
	var party = []pokemonbattleadvisor.Pokemon{
		{
			Name: "weavile", Types: []string{"dark", "ice"},
			Abilities: []string{"pressure"},
			Moves: []pokemonbattleadvisor.Move{
				{Name: "poison-jab", Type: "poison", Damage: 80.0, Accuracy: 1.00},
				{Name: "false-swipe", Type: "normal", Damage: 40.0, Accuracy: 1.00},
				{Name: "hail", Type: "ice", Damage: 0.0, Accuracy: 100.0},
				{Name: "blizzard", Type: "ice", Damage: 110.0, Accuracy: 0.70},
			},
		},
		{
			Name: "clefable", Types: []string{"fairy"},
			Abilities: []string{"unaware"},
			Moves: []pokemonbattleadvisor.Move{
				{Name: "moonblast", Type: "fairy", Damage: 95.0, Accuracy: 1.00},
				{Name: "flash", Type: "normal", Damage: 0.0, Accuracy: 1.0},
				{Name: "flamethrower", Type: "fire", Damage: 90.0, Accuracy: 1.00},
				{Name: "double-slap", Type: "normal", Damage: 15.0, Accuracy: 0.85},
			},
		},
		{
			Name: "azumarill", Types: []string{"water", "fairy"},
			Abilities: []string{"huge-power"},
			Moves: []pokemonbattleadvisor.Move{
				{Name: "ice-beam", Type: "water", Damage: 90.0, Accuracy: 1.00},
				{Name: "play-rough", Type: "fairy", Damage: 90.0, Accuracy: 0.90},
				{Name: "surf", Type: "water", Damage: 90.0, Accuracy: 1.00},
				{Name: "hydro-pump", Type: "water", Damage: 110.0, Accuracy: 0.80},
			},
		},
		{
			Name: "luxray", Types: []string{"electric"},
			Abilities: []string{"rivalry"},
			Moves: []pokemonbattleadvisor.Move{
				{Name: "thunderbolt", Type: "electric", Damage: 90.0, Accuracy: 1.00},
				{Name: "crunch", Type: "dark", Damage: 80.0, Accuracy: 1.00},
				{Name: "flash", Type: "normal", Damage: 0.0, Accuracy: 1.00},
				{Name: "discharge", Type: "electric", Damage: 80.0, Accuracy: 1.00},
			},
		},
		{
			Name: "ludicolo", Types: []string{"water", "grass"},
			Abilities: []string{"swift-swim"},
			Moves: []pokemonbattleadvisor.Move{
				{Name: "dive", Type: "water", Damage: 80.0, Accuracy: 1.00},
				{Name: "surf", Type: "water", Damage: 90.0, Accuracy: 1.00},
				{Name: "giga-drain", Type: "grass", Damage: 75.0, Accuracy: 1.00},
				{Name: "energy-ball", Type: "grass", Damage: 90.0, Accuracy: 1.00},
			},
		},
		{
			Name: "sceptile", Types: []string{"grass"},
			Abilities: []string{"overgrow"},
			Moves: []pokemonbattleadvisor.Move{
				{Name: "cut", Type: "normal", Damage: 50.0, Accuracy: 0.95},
				{Name: "dig", Type: "ground", Damage: 80.0, Accuracy: 1.00},
				{Name: "energy-ball", Type: "grass", Damage: 90.0, Accuracy: 1.00},
				{Name: "giga-drain", Type: "grass", Damage: 75.0, Accuracy: 1.00},
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
