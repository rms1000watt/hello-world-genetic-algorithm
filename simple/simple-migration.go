package simple

import (
	"fmt"

	"github.com/rms1000watt/hello-world-genetic-algorithm/ga"
)

type SimpleMigration struct {
	Count       int
	Max         int
	MigrationCh chan SimpleIndividual
}

func NewSimpleMigration(size int) ga.Migration {
	return SimpleMigration{
		Max:         size,
		MigrationCh: make(chan SimpleIndividual, size),
	}
}

func (m SimpleMigration) Pop() (ga.Individual, bool) {
	if m.Count == 0 {
		fmt.Println("Migration empty")
		return nil, false
	}

	m.Count--
	return <-m.MigrationCh, true
}

func (m SimpleMigration) Push(ind ga.Individual) bool {
	simpleInd, ok := ind.(SimpleIndividual)
	if !ok {
		fmt.Println("Not SimpleIndividual.. not pushing to migration..")
		return false
	}

	if m.Count == m.Max {
		fmt.Println("Migration full")
		return false
	}

	m.MigrationCh <- simpleInd
	return true
}

func (m SimpleMigration) Flush() {
	cnt := int(0)
	for {
		select {
		case <-m.MigrationCh:
			cnt++
		default:
			goto here
		}
	}
here:

	fmt.Println("Flushed from migration:", cnt)
}
