package simple

import (
	"fmt"

	"github.com/rms1000watt/hello-world-genetic-algorithm/ga"
)

type SimpleMigration chan SimpleIndividual

func NewSimpleMigration(size int) ga.Migration {
	return SimpleMigration(make(chan SimpleIndividual, size))
}

func (m SimpleMigration) Pop() (ga.Individual, bool) {
	if len(m) == 0 {
		return nil, false
	}

	return <-m, true
}

func (m SimpleMigration) Push(ind ga.Individual) bool {
	simpleInd, ok := ind.(SimpleIndividual)
	if !ok {
		fmt.Println("Not SimpleIndividual.. not pushing to migration..")
		return false
	}

	if len(m) == cap(m) {
		return false
	}

	m <- simpleInd
	return true
}

func (m SimpleMigration) Flush() {
	cnt := int(0)
	for {
		select {
		case <-m:
			cnt++
		default:
			goto here
		}
	}
here:

	fmt.Println("Flushed from migration:", cnt)
}

func (m SimpleMigration) Length() int {
	return len(m)
}
