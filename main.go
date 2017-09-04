package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/rms1000watt/hello-world-genetic-algorithm/ga"
	"github.com/rms1000watt/hello-world-genetic-algorithm/simple"
)

const (
	populationSize = int(1e4)
	retainSize     = int(8e3) // Must be >= population size
	mutationFactor = int(90)  // Must be between 0 and 100
	iterations     = int(1e3)
)

var wg sync.WaitGroup

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func main() {
	migration := simple.NewSimpleMigration(populationSize * 2)

	population := simple.NewSimplePopulation(populationSize)
	evolver := simple.NewSimpleEvolver()
	Run(population, evolver, migration)

	migration.Flush()

	// population = simple.NewSimplePopulation(populationSize)
	// evolver = concurrent.NewConcurrentEvolver()
	// RunConcurrent(population, evolver, migration)
}

func RunConcurrent(pop ga.Population, evolver ga.Evolver, migration ga.Migration) {
	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		go Run(pop, evolver, migration)
	}
}

func Run(pop ga.Population, evolver ga.Evolver, migration ga.Migration) {
	now := time.Now()
	fmt.Println("Population Size:", populationSize)
	fmt.Println("Iterations:", iterations)
	fmt.Println("Start Grade:", pop.Grade())

	for i := 0; i < iterations; i++ {
		pop = evolver.Evolve(pop, retainSize, mutationFactor, migration)
	}

	fmt.Println("Done Grade:", pop.Grade())
	fmt.Println("Run Time:", time.Since(now))
}
