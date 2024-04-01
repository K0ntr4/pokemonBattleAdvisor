package main

func main() {
	team := GetRandomTeam(0, 493)
	for _, pokemon := range team {
		println(pokemon.Name)
	}
}
