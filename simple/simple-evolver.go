package simple

import (
	"math/rand"
)

type SimpleEvolver struct{}

func (e SimpleEvolver) Evolve(pop Population, retain, mutationFactor int) Population {
	// fmt.Println("Pop:", pop)

	// Get the population size for later
	popSize := pop.Length()

	// Sort population by best individuals
	pop = pop.Sort()
	// fmt.Println("Sorted:", pop)

	// Retain only the best individuals (based on retain parameter)
	pop = pop.Best(retain)
	// fmt.Println("Best:", pop)

	// Randomly add new individuals to next population set
	newPop := NewSimplePopulation(popSize - retain - rand.Intn(popSize-retain))
	// fmt.Println("New Pop:", newPop)

	// Merge the random individuals
	pop = pop.Merge(newPop)
	// fmt.Println("Merged:", pop)

	// Randomly mutate some individuals
	pop = pop.Mutate(mutationFactor)
	// fmt.Println("Mutated:", pop)

	// Crossbreed population for remaining space
	for pop.Length() < popSize {
		mInd := rand.Intn(pop.Length())
		dInd := rand.Intn(pop.Length())
		if mInd == dInd {
			continue
		}

		mom := pop.At(mInd)
		dad := pop.At(dInd)

		child := mom.Breed(dad)
		pop = pop.Add(child)
	}
	// fmt.Println("Bread:", pop)

	return pop
}

func NewSimpleEvolver() Evolver {
	return SimpleEvolver{}
}
