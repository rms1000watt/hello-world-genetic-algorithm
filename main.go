package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rms1000watt/hello-world-genetic-algorithm/simple"
)

// Dynamic retainSize and mutationFactor?

const (
	populationSize = int(10)
	retainSize     = int(8)
	mutationFactor = int(90)
	iterations     = int(10000)
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	population := simple.NewSimplePopulation(populationSize)
	evolver := simple.NewSimpleEvolver()

	fmt.Println("Start:", population, population.Grade())
	population = Run(population, evolver)
	fmt.Println("Done:", population, population.Grade())
}

func Run(pop simple.Population, evolver simple.Evolver) simple.Population {
	for i := 0; i < iterations; i++ {
		pop = evolver.Evolve(pop, retainSize, mutationFactor)
	}
	return pop
}
