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
	migrationCh := make(chan ga.Individual, populationSize*2)
	population := simple.NewSimplePopulation(populationSize)
	evolver := simple.NewSimpleEvolver()
	Run(population, evolver, migrationCh)

	flushMigrationCh(migrationCh)

	population = simple.NewSimplePopulation(populationSize)
	evolver = concurrent.NewConcurrentEvolver()
	RunConcurrent(population, evolver, migrationCh)
}

func RunConcurrent(pop ga.Population, evolver ga.Evolver, migrationCh chan ga.Individual) {
	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		go Run(pop, evolver, migrationCh)
	}
}

func Run(pop ga.Population, evolver ga.Evolver, migrationCh chan ga.Individual) {
	now := time.Now()
	fmt.Println("Population Size:", populationSize)
	fmt.Println("Iterations:", iterations)
	fmt.Println("Start Grade:", pop.Grade())

	for i := 0; i < iterations; i++ {
		pop = evolver.Evolve(pop, retainSize, mutationFactor, migrationCh)
	}

	fmt.Println("Done Grade:", pop.Grade())
	fmt.Println("Run Time:", time.Since(now))
}

func flushMigrationCh(migrationCh chan ga.Individual) {
	cnt := int(0)
	for {
		select {
		case <-migrationCh:
			fmt.Println("cnt++")
			cnt++
		default:
			fmt.Println("goto here")
			goto here
		}
	}
here:

	fmt.Println("Flushed from migration channel:", cnt)
}
