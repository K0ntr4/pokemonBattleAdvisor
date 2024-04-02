package pokemon_battle_advisor

func SetDefaultBounds(bounds *[]int) []int {
	// Set default bounds if not provided
	switch len(*bounds) {
	case 0:
		*bounds = []int{0, 493} // Default to gen 4
	case 1:
		*bounds = append(*bounds, 493)
	}
	return *bounds
}
