# Pokemon Battle Advisor

**Description**

The Pokemon Battle Advisor is a tool designed to help you make better decisions during battles. It provides insights and recommendations based on your party members and the enemy Pokemon you're facing.

**Features**

- **Programmatic Usage**: Easily integrate the Pokemon Battle Advisor into your own projects.

- **Customizable Party**: Define your party members with their names, types, abilities, and moves.

- **Automatic Enemy Detection**: Automatically identify the enemy Pokemon by taking a screenshot of the battle.

- **Multi-Hit Moves Handling**: Takes into account multi-hit moves by calculating average damage, accuracy, and overall effectiveness.

- **Dynamic Recommendations**: Get recommendations on the best strategy to use against the enemy Pokemon, including whether to switch party members and which move to use.

**Factors Considered for Determining Best Move and Party Member:**

- **Move Effectiveness**: Evaluates the effectiveness of each move against the enemy Pokemon based on their type and weaknesses.

- **Damage Calculation**: Considers the damage output of each move, taking into account factors such as base power, type effectiveness, and the attacker's stats.

- **Accuracy**: Accounts for the accuracy of each move to determine the likelihood of it hitting the target.

- **Party Member Attributes**: Analyzes the attributes of each party member, including types, abilities, and moves, to determine the best party member for the battle.

## Installation

To install the Pokemon Battle Advisor, follow these steps:

1. Clone the repository: `git clone https://github.com/K0ntr4/pokemonBattleAdvisor.git`
2. Navigate to the project directory: `cd pokemon-battle-advisor`
3. Install dependencies: `go mod tidy`

**Usage**

The Pokemon Battle Advisor provides a convenient way to strategize and optimize your battle decisions programmatically. Here's how you can use it:

1. **Define Your Party:**

   ```go
    func main() {
        var party = []pokemonbattleadvisor.Pokemon{
		    pokemonbattleadvisor.GetPartyPokemon("weavile", []string{"pressure"}, []string{"poison-jab", "false-swipe", "hail", "blizzard"}),
        }
        // Add more party members as needed
    }
   ```

   Create an array of your own party members, each with their own name, types, abilities, and moves. Customize your party based on your preferences and strategy.

2. **Automatically Determine Enemy Pokemon:**

   ```go
   func getEnemyPokemon() (enemy pokemonbattleadvisor.Pokemon) {
       // Capture a screenshot of the battle
       screenshot, err = pokemonbattleadvisor.TakeScreenshot(0, 1250, 450, 1600, 800) // Adjust display and coordinates as needed
       if err != nil {
           return
       }

       // Classify the screenshot to identify the enemy Pokemon
       classifierResult, err := pokemonbattleadvisor.Classify(screenshot)
       if err != nil {
           // Handle error or wait for classification to complete
       }

       // Retrieve the enemy Pokemon using the classification result
       enemy, err = pokemonbattleadvisor.PokemonByName(classifierResult)
       if err != nil {
           panic(err)
       }

       // Print details of the enemy Pokemon
       pokemonbattleadvisor.PrintHelperStructsPokemon(&enemy)

       return enemy
   }
   ```

   Utilize the `getEnemyPokemon()` function to automatically detect the enemy Pokemon in the battle by taking a screenshot and classifying it. This method is optimized for the default resolutions of the game Pokemon Revolution. For other resolutions or different games, consider using the `PokemonByName` function with the name of the enemy Pokemon directly for improved accuracy.

3. **Assess the Best Strategy:**

   ```go
   partyMember, move, shouldSwitch := pokemonbattleadvisor.BestPokemonMoveAndShouldSwitch(&party, &enemy)
   if shouldSwitch {
       fmt.Println("Should switch")
       fmt.Printf("Best party member: %s\n", party[partyMember].Name)
       fmt.Printf("Best move: %s\n", party[partyMember].Moves[move].Name)
   } else {
       fmt.Println("Should not switch")
       fmt.Printf("Best move: %s\n", party[0].Moves[move].Name)
   }
   ```

   Use the `BestPokemonMoveAndShouldSwitch()` function to determine the best party member and move to use against the enemy Pokemon. This function evaluates various factors and provides recommendations based on the current battle scenario.

By following these steps, you can effectively utilize the Pokemon Battle Advisor to enhance your strategic decision-making and improve your chances of success in battles. Adjust the party composition and strategies as needed to suit your preferences and objectives.

## Contributing

Contributions to the Pokemon Battle Advisor project are welcome! If you find any bugs or have suggestions for improvements, feel free to open an issue or submit a pull request.

## License

The Pokemon Battle Advisor project is licensed under the MIT License. See the [LICENSE](https://github.com/yourusername/pokemon-battle-advisor/blob/main/LICENSE) file for more details.

## Acknowledgements

The Pokemon Battle Advisor project utilizes several external libraries and resources:

- [hfapigo/v2](https://pkg.go.dev/github.com/Kardbord/hfapigo/v2@v2.1.0): Provides an API interface to Hugging Face for automatic Pokemon detection.
- [imjeffhi/pokemon_classifier](https://huggingface.co/imjeffhi/pokemon_classifier): Utilizes a Hugging Face model for Pokemon classification.
- [sajari/fuzzy](https://pkg.go.dev/github.com/sajari/fuzzy@v1.0.0): Implements fuzzy search functionality for improved user experience.
- [kbinani/screenshot](https://pkg.go.dev/github.com/kbinani/screenshot@v0.0.0-20230812210009-b87d31814237): Enables capturing screenshots for automatic enemy detection.
- [olekukonko/typewriter](https://pkg.go.dev/github.com/olekukonko/tablewriter@v0.0.5): Provides table formatting functionality for improved output presentation.

Special thanks to the developers and maintainers of these libraries for their contributions to the Pokemon Battle Advisor project. Their work greatly enhances the functionality and usability of the tool.