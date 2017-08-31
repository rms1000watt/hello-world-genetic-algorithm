package simple

import (
	"math/rand"
)

type SimpleEvolver struct{}

func (e SimpleEvolver) Evolve(pop Population, retain, mutationFactor int) Population {
	// Get the population size for later
	popSize := pop.Length()

	// Sort population by best individuals
	pop = pop.Sort()

	// Retain only the best individuals (based on retain parameter)
	pop = pop.Best(retain)

	// Randomly add new individuals to next population set
	newPop := NewSimplePopulation(popSize - retain - rand.Intn(popSize-retain))

	// Merge the random individuals
	pop = pop.Merge(newPop)

	// Randomly mutate some individuals
	pop = pop.Mutate(mutationFactor)

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

	return pop
}

func NewSimpleEvolver() Evolver {
	return SimpleEvolver{}
}
