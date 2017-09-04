package concurrent

import (
	"math/rand"

	"github.com/rms1000watt/hello-world-genetic-algorithm/ga"
)

type ConcurrentEvolver struct{}

func NewConcurrentEvolver() ga.Evolver {
	return ConcurrentEvolver{}
}

func (e ConcurrentEvolver) Evolve(pop ga.Population, retain, mutationFactor int, migration ga.Migration) ga.Population {
	// Get the population size for later
	popSize := pop.Length()

	// Sort population by best individuals
	pop = pop.Sort()

	// Retain only the best individuals (based on retain parameter)
	pop = pop.Best(retain)

	// Emigrate some of the best
	for i := 0; i < rand.Intn(3); i++ {
		pop, ind := pop.Pop()
		if ok := migration.Push(ind); !ok {
			break
		}
	}

	// Immigrate some of the best

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
		pop = pop.Push(child)
	}

	return pop
}
