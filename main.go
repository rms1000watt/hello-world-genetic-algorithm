package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"

	"github.com/rms1000watt/hello-world-genetic-algorithm/concurrent"
	"github.com/rms1000watt/hello-world-genetic-algorithm/ga"
	"github.com/rms1000watt/hello-world-genetic-algorithm/simple"
)

const (
	populationSize  = int(1e3)
	retainSize      = int(8e2) // Must be <= population size
	mutationFactor  = int(10)  // Must be between 0 and 100. Greater the factor, greater the chance for mutation
	migrationFactor = int(10)  // Must be between 0 and 100. Greater the factor, greater the chance for migration
	iterations      = int(3e3)
	migrationSize   = int(populationSize * 3)
)

var wg sync.WaitGroup

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func main() {
	// Create a migration so individuals can go between populations
	// Populations are each evolved in their own goroutines
	migration := simple.NewSimpleMigration(migrationSize)

	// SimpleIndividuals are just random ints. SimplePopulations contain
	// simpleIndividuals
	population := simple.NewSimplePopulation(populationSize)
	evolver := simple.NewSimpleEvolver()

	// Run a single evolution job
	wg.Add(1)
	go Run(population, evolver, migration)
	wg.Wait()

	// ConcurrentEvolver uses the migration to pass individuals between populations
	evolver = concurrent.NewConcurrentEvolver()

	// Run a concurrent evolution job
	RunSimpleConcurrent(populationSize, evolver, migration)
	wg.Wait()
}

func RunSimpleConcurrent(populationSize int, evolver ga.Evolver, migration ga.Migration) {
	for i := 0; i < runtime.GOMAXPROCS(0)-1; i++ {
		wg.Add(1)
		go Run(simple.NewSimplePopulation(populationSize), evolver, migration)
	}
}

func Run(pop ga.Population, evolver ga.Evolver, migration ga.Migration) {
	now := time.Now()
	runID := rand.Intn(1000)
	fmt.Println(runID, "Population Size:", populationSize)
	fmt.Println(runID, "Iterations:", iterations)
	fmt.Println(runID, "Start Grade:", pop.Grade())

	for i := 0; i < iterations; i++ {
		pop = evolver.Evolve(pop, retainSize, mutationFactor, migrationFactor, migration)
	}

	fmt.Println(runID, "Done Grade:", pop.Grade())
	fmt.Println(runID, "Unused in Migration:", migration.Length())
	fmt.Println(runID, "Run Time:", time.Since(now))

	wg.Done()
}
