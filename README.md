# Pokemon Battle Advisor

The Pokemon Battle Advisor is a GoLang project designed to provide recommendations on which Pokemon and moves to use against a specific target Pokemon in a battle. It takes into account various factors such as move types, Pokemon types, and abilities to determine the most effective strategy.

## Features

- **Type Effectiveness Calculation:** Calculates the effectiveness of a move against a target Pokemon based on type matchups.
- **Ability Consideration:** Takes into account the abilities of both the attacking and defending Pokemon to adjust the effectiveness of moves.
- **Integration with PokeAPI:** Utilizes the PokeAPI to retrieve information about Pokemon, including their types and abilities.

## Installation

To install the Pokemon Battle Advisor, follow these steps:

1. Clone the repository: `git clone https://github.com/yourusername/pokemon-battle-advisor.git`
2. Navigate to the project directory: `cd pokemon-battle-advisor`
3. Install dependencies: `go mod tidy`

## Usage

To use the Pokemon Battle Advisor, you can interact with it programmatically by importing the `helperStructs` package and calling its functions. Below is a basic example:

```go
import (
    "fmt"
    "github.com/yourusername/pokemon-battle-advisor/helperStructs"
)

func main() {
    // Create a Move instance
    move := helperStructs.Move{Name: "tackle", Type: "normal"}

    // Create a target Pokemon instance
    targetPokemon := helperStructs.Pokemon{
        Abilities: []string{"levitate"},
        Moves:     []helperStructs.Move{},
        Types:     []string{"ghost"},
        Name:      "Gengar",
    }

    // Determine the effectiveness of the move against the target Pokemon
    effectiveness := move.EffectivenessAgainst(&targetPokemon, &[]string{})

    fmt.Printf("Effectiveness of %s move against %s: %f\n", move.Name, targetPokemon.Name, effectiveness)
}
```

## Contributing

Contributions to the Pokemon Battle Advisor project are welcome! If you find any bugs or have suggestions for improvements, feel free to open an issue or submit a pull request.

## License

The Pokemon Battle Advisor project is licensed under the MIT License. See the [LICENSE](https://github.com/yourusername/pokemon-battle-advisor/blob/main/LICENSE) file for more details.

## Acknowledgements

The Pokemon Battle Advisor project utilizes the [PokeAPI](https://pokeapi.co/) for retrieving information about Pokemon. Special thanks to the developers and maintainers of the PokeAPI for their work in providing this valuable resource.
