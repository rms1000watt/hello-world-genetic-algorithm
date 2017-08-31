package simple

import (
	"fmt"
	"math/rand"
)

const (
	randMaxFloat = 1e6 // Must be greater than 1e3
	randMaxInt   = int(randMaxFloat)
)

type SimpleIndividual int

func NewSimpleIndividual() Individual {
	return SimpleIndividual(rand.Intn(randMaxInt))
}

func (i SimpleIndividual) Fit() int {
	return int(i) / int(randMaxFloat*1e-2)
}

func (i SimpleIndividual) Value() interface{} {
	return int(i)
}

func (i SimpleIndividual) Breed(ind Individual) Individual {
	inSimpleInd, ok := ind.(SimpleIndividual)
	if !ok {
		fmt.Println("BREED: individual not SimpleIndividual.. using NewSimpleIndividual")
		inSimpleInd = NewSimpleIndividual().(SimpleIndividual)
	}

	return SimpleIndividual((int(i) + int(inSimpleInd)) / 2)
}
