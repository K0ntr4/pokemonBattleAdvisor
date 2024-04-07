package pokemonbattleadvisor

var (
	typeEffects = map[string]map[string]float64{
		"normal":   {"fighting": 2.0, "ghost": 0},
		"fire":     {"fire": 0.5, "water": 2.0, "grass": 0.5, "ice": 0.5, "ground": 2.0, "bug": 0.5, "rock": 2.0, "steel": 0.5, "fairy": 0.5},
		"water":    {"fire": 0.5, "water": 0.5, "electric": 2.0, "grass": 2.0, "ice": 0.5, "steel": 0.5},
		"electric": {"electric": 0.5, "ground": 2.0, "flying": 0.5, "steel": 0.5},
		"grass":    {"fire": 2.0, "water": 0.5, "grass": 0.5, "ice": 2.0, "poison": 2.0, "ground": 0.5, "flying": 2.0, "bug": 2.0},
		"ice":      {"fire": 2.0, "ice": 0.5, "fighting": 2.0, "rock": 2.0, "steel": 2.0},
		"fighting": {"flying": 2.0, "psychic": 2.0, "bug": 0.5, "rock": 0.5, "dark": 0.5, "fairy": 2.0},
		"poison":   {"grass": 0.5, "fighting": 0.5, "poison": 0.5, "ground": 2.0, "psychic": 2.0, "fairy": 0.5},
		"ground":   {"water": 2.0, "electric": 0.0, "grass": 2.0, "ice": 2.0, "poison": 0.5, "rock": 0.5},
		"flying":   {"electric": 2.0, "grass": 0.5, "ice": 2.0, "fighting": 0.5, "ground": 0.0, "bug": 0.5, "rock": 2.0},
		"psychic":  {"fighting": 0.5, "psychic": 0.5, "bug": 2.0, "ghost": 2.0, "dark": 2.0},
		"bug":      {"fire": 2.0, "grass": 0.5, "fighting": 0.5, "ground": 0.5, "flying": 2.0, "rock": 2.0},
		"rock":     {"normal": 0.5, "fire": 0.5, "water": 2.0, "grass": 2.0, "fighting": 2.0, "poison": 0.5, "ground": 2.0, "flying": 0.5, "steel": 2.0},
		"ghost":    {"normal": 0.0, "fighting": 0.0, "poison": 0.5, "bug": 0.5, "ghost": 2.0, "dark": 2.0},
		"dragon":   {"fire": 0.5, "water": 0.5, "electric": 0.5, "grass": 0.5, "ice": 2.0, "dragon": 2.0, "fairy": 2.0},
		"dark":     {"fighting": 2.0, "psychic": 0.0, "bug": 2.0, "ghost": 0.5, "dark": 0.5, "fairy": 2.0},
		"steel":    {"normal": 0.5, "fire": 2.0, "grass": 0.5, "ice": 0.5, "fighting": 2.0, "poison": 0.0, "ground": 2.0, "flying": 0.5, "psychic": 0.5, "bug": 0.5, "rock": 0.5, "dragon": 0.5, "steel": 0.5, "fairy": 0.5},
		"fairy":    {"fighting": 0.5, "poison": 2.0, "bug": 0.5, "dragon": 0.0, "dark": 0.5, "steel": 2.0},
	}

	abilitiesEffects = map[string]map[string]float64{
		"dry-skin":     {"water": 0.0, "fire": 1.25},
		"flash-fire":   {"fire": 0.0},
		"heatproof":    {"fire": 0.5},
		"levitate":     {"ground": 0.0},
		"thick-fat":    {"fire": 0.5, "ice": 0.5},
		"volt-absorb":  {"electric": 0.0},
		"motor-drive":  {"electric": 0.0},
		"water-absorb": {"water": 0.0},
		"sap-sipper":   {"grass": 0.0},
	}
)

type Pokemon struct {
	Name      string   `json:"name"`
	Types     []string `json:"type"`
	Abilities []string `json:"ability"`
	Moves     []Move   `json:"moves"`
}

type Move struct {
	Rank     int     `json:"rank"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Damage   float64 `json:"damage"`
	Accuracy float64 `json:"accuracy"`
}

func ApplySpecialAbilities(attackerAbilities *[]string, moveType *string, scrappy *bool) {
	for _, a := range *attackerAbilities {
		switch a {
		case "normalize": // Normalize changes all moves to normal type
			*moveType = "normal"
		case "scrappy": // Scrappy allows normal and fighting moves to hit ghost types
			*scrappy = true
		}
	}
}

func CalculateTypeEffectiveness(types *[]string, moveType *string, scrappy bool, result *float64) {
	for _, t := range *types {
		multiplier, exists := typeEffects[t][*moveType]
		if !exists {
			continue
		}
		if scrappy && t == "ghost" && (*moveType == "normal" || *moveType == "fighting") { // Scrappy allows normal and fighting moves to hit ghost types
			multiplier = 1.0
		}
		*result *= multiplier
	}
}

func ApplyEnemyPokemonAbilities(abilities *[]string, moveType *string, result *float64) {
	for _, a := range *abilities {
		if a == "wonder-guard" && *result <= 1.0 { // Wonder Guard only allows super effective moves
			*result = 0.0
			return
		}
		if (a == "filter" || a == "solid-rock") && *result > 1.0 { // Filter and Solid Rock reduce super effective moves
			*result *= 0.75
		}
		ability, exists := abilitiesEffects[a]
		if !exists {
			continue
		}
		if multiplier, exists := ability[*moveType]; exists {
			*result *= multiplier
		}
	}
}

func (m Move) EffectivenessAgainst(pokemon *Pokemon, attackerAbilities *[]string) (result float64) {
	result = 1.0
	var moveType = m.Type
	var scrappy = false

	ApplySpecialAbilities(attackerAbilities, &moveType, &scrappy)

	CalculateTypeEffectiveness(&pokemon.Types, &moveType, scrappy, &result)

	ApplyEnemyPokemonAbilities(&pokemon.Abilities, &moveType, &result)

	return result
}
