package concurrent

import (
	"github.com/rms1000watt/hello-world-genetic-algorithm/ga"
)

type ConcurrentEvolver struct{}

func NewConcurrentEvolver() ga.Evolver {
	return ConcurrentEvolver{}
}

func (e ConcurrentEvolver) Evolve(pop ga.Population, retain, mutationFactor int, migrationCh chan ga.Individual) ga.Population {
	return pop
}
