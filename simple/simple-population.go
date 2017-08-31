package simple

import (
	"fmt"
	"math/rand"
	"sort"
)

type SimplePopulation []Individual

func NewSimplePopulation(size int) Population {
	var population SimplePopulation
	for i := 0; i < size; i++ {
		population = append(population, NewSimpleIndividual())
	}

	return population
}

func (p SimplePopulation) Grade() int {
	var grade int
	for _, i := range p {
		grade += i.Fit()
	}

	return grade / int(len(p))
}

func (p SimplePopulation) Sort() Population {
	var ints []int
	for _, i := range p {
		ints = append(ints, i.Value().(int))
	}

	sort.Ints(ints)

	var newPop SimplePopulation
	for _, i := range ints {
		newPop = append(newPop, SimpleIndividual(i))
	}

	return newPop
}

func (p SimplePopulation) Best(size int) Population {
	if size > len(p) {
		fmt.Println("BEST: size > len(population).. returning all..")
		return p
	}

	var newPop SimplePopulation
	for i := len(p) - 1; i > len(p)-size-1; i-- {
		// for i := 0; i < size; i-- {
		newPop = append(newPop, p[i])
	}

	return newPop
}

func (p SimplePopulation) Merge(pop Population) Population {
	for _, ind := range pop.(SimplePopulation) {
		p = append(p, ind)
	}

	return p
}

func (p SimplePopulation) Length() int {
	return len(p)
}

func (p SimplePopulation) Mutate(mutationFactor int) Population {
	var newPop SimplePopulation
	for _, ind := range p {
		if rand.Intn(100) > mutationFactor {
			newPop = append(newPop, NewSimpleIndividual())
			continue
		}
		newPop = append(newPop, ind)
	}
	return newPop
}

func (p SimplePopulation) At(i int) Individual {
	return p[i]
}

func (p SimplePopulation) Add(ind Individual) Population {
	simpleInd, ok := ind.(SimpleIndividual)
	if !ok {
		fmt.Println("ADD: Indivdual not SimpleIndividual.. appending NewSimpleIndividual")
		return append(p, NewSimpleIndividual())
	}

	return append(p, simpleInd)
}
