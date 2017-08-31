package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rms1000watt/hello-world-genetic-algorithm/simple"
)

const (
	populationSize = int(5)
	retainSize     = int(2)
	mutationFactor = int(75)
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	population := simple.NewSimplePopulation(populationSize)
	fmt.Println(population)

	evolver := simple.NewSimpleEvolver()
	population = evolver.Evolve(population, retainSize, mutationFactor)
}
