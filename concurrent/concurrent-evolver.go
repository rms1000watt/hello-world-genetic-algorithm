package concurrent

import (
	"math/rand"

	"github.com/rms1000watt/hello-world-genetic-algorithm/ga"
	"github.com/rms1000watt/hello-world-genetic-algorithm/simple"
)

type ConcurrentEvolver struct{}

func NewConcurrentEvolver() ga.Evolver {
	return ConcurrentEvolver{}
}

func (e ConcurrentEvolver) Evolve(pop ga.Population, retain, mutationFactor, migrationFactor int, migration ga.Migration) ga.Population {
	// Get the population size for later
	popSize := pop.Length()

	// Sort population by best individuals
	pop = pop.Sort()

	// Retain only the best individuals (based on retain parameter)
	pop = pop.Best(retain)

	// Emigrate some of the best
	if rand.Intn(100) < migrationFactor {
		// TODO: Tune this..
		emigrateSize := rand.Intn(pop.Length() / 1)
		pop = pop.Emigrate(emigrateSize, migration)
	}

	// Immigrate some of the best
	pop = pop.Immigrate(rand.Intn(popSize-pop.Length()), migration)

	// Randomly add new individuals to next population set
	newPop := simple.NewSimplePopulation(popSize - pop.Length() - rand.Intn(popSize-pop.Length()))

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

	// TODO: Grade the population.. if high grade.. kill early?

	return pop
}
