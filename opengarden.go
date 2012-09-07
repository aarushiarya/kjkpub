package main

import (
	"fmt"
	"math/rand"
)

var population = [...]int{18897109, 12828837, 9461105, 6371773, 5965343, 5946800, 5582170, 5564635, 5268860, 4552402, 4335391, 4296250, 4224851, 4192887, 3439809, 3279833, 3095313, 2812896, 2783243, 2710489, 2543482, 2356285, 2226009, 2149127, 2142508, 2134411}

const desired = 100000000

func isBitSet(n int, bit int) bool {
	return n&(1<<uint(bit)) != 0
}

func setBit(n int, bit int) int {
	return n | (1 << uint(bit))
}

func pickRandom(picked *int) int {
	bf := *picked
	for {
		n := rand.Intn(len(population))
		if isBitSet(bf, n) {
			continue // has already been picked
		}
		*picked = setBit(bf, n) // mark as picked
		return int(n)
	}
	return 0
}

func tryRandom(solution []int) int {
	var picked int = 0
	remaining := desired
	for {
		pickedIdx := pickRandom(&picked)
		remaining -= population[pickedIdx]
		if 0 == remaining {
			return picked
		}
		if remaining < 0 {
			return 0
		}
	}
	return 0
}

func printSolution(solution int) {
	fmt.Print("Solution: ")
	for i := 0; i < len(population); i++ {
		if isBitSet(solution, i) {
			fmt.Printf("%d, ", population[i])
		}
	}
	fmt.Print("\n")
}

func random() {
	solution := make([]int, len(population))
	rounds := 0
	for {
		rounds++
		if rounds%40000 == 0 {
			fmt.Printf("random(): %d rounds\n", rounds)
		}
		if rounds%10000000 == 0 {
			break
		}
		if solution := tryRandom(solution); solution != 0 {
			printSolution(solution)
			return
		}
	}
}

func main() {
	random()
}