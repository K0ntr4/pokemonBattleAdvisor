package helperStructs

var allTypes = map[string]map[string]float64{
	"normal":   {"fighting": 2.0, "ghost": 0},
	"fire":     {"fire": 0.5, "water": 2.0, "grass": 0.5, "ice": 0.5, "ground": 2.0, "bug": 0.5, "rock": 2.0, "steel": 0.5, "fairy": 0.5},
	"water":    {"fire": 0.5, "water": 0.5, "electric": 2.0, "grass": 2.0, "ice": 0.5, "steel": 0.5},
	"electric": {"electric": 0.5, "ground": 2.0, "flying": 0.5, "steel": 0.5},
	"grass":    {"fire": 2.0, "water": 0.5, "grass": 0.5, "ice": 2.0, "poison": 2.0, "ground": 0.5, "flying": 2.0, "bug": 2.0},
	"ice":      {"fire": 2.0, "ice": 0.5, "fighting": 2.0, "rock": 2.0, "steel": 2.0},
	"fighting": {"flying": 2.0, "psychic": 2.0, "bug": 0.5, "rock": 0.5, "dark": 0.5, "fairy": 2.0},
	"poison":   {"grass": 0.5, "fighting": 0.5, "poison": 0.5, "ground": 2.0, "psychic": 2.0, "fairy": 0.5},
	"ground":   {"water": 2.0, "electric": 0.0, "grass": 2.0, "ice": 2.0, "poison": 1 / 2.0, "rock": 0.5},
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

// filter, normalize, scrappy, solid-rock, wonder-guard special cases where the ability affects the effectiveness of the move directly

var abilities = map[string]map[string]float64{
	"dry-skin":     {"water": 0.0, "fire": 1.25},
	"flash-fire":   {"fire": 0.0},
	"heatproof":    {"fire": 0.5},
	"levitate":     {"ground": 0.0},
	"thick-fat":    {"fire": 0.5, "ice": 0.5},
	"volt-absorb":  {"electric": 0.0},
	"motor-drive":  {"electric": 0.0},
	"water-absorb": {"water": 0.0},
}

type Pokemon struct {
	Abilities []string `json:"ability"`
	Moves     []Move   `json:"moves"`
	Types     []string `json:"type"`
	Name      string   `json:"name"`
}

type Move struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (m Move) EffectivenessAgainst(pokemon Pokemon) float64 {
	var res float64 = 1.0
	for _, t := range pokemon.Types {
		multiplier, exists := allTypes[t][m.Type]
		if !exists {
			continue
		}
		res *= multiplier
	}
	for _, a := range pokemon.Abilities {
		ability, exists := abilities[a]
		if !exists {
			continue
		}
		for t, multiplier := range ability {
			if t == m.Type {
				res *= multiplier
			}
		}
	}
	return res
}
