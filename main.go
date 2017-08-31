package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rms1000watt/hello-world-genetic-algorithm/simple"
)

const (
	populationSize = int(1e4)
	retainSize     = int(8e3) // Must be >= population size
	mutationFactor = int(90)  // Must be between 0 and 100
	iterations     = int(1e3)
)

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func main() {
	population := simple.NewSimplePopulation(populationSize)
	evolver := simple.NewSimpleEvolver()

	Run(population, evolver)
}

func Run(pop simple.Population, evolver simple.Evolver) {
	now := time.Now()
	fmt.Println("Population Size:", populationSize)
	fmt.Println("Iterations:", iterations)
	fmt.Println("Start Grade:", pop.Grade())

	for i := 0; i < iterations; i++ {
		pop = evolver.Evolve(pop, retainSize, mutationFactor)
	}

	fmt.Println("Done Grade:", pop.Grade())
	fmt.Println("Run Time:", time.Since(now))
}
