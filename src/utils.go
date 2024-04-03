package pokemonbattleadvisor

import (
	"crypto/rand"
	"math/big"
)

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

func GetRandomIndex[T any](list []T) int {
	result, err := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
	if err != nil {
		panic(err)
	}
	return int(result.Int64())
}
