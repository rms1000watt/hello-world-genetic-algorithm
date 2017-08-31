package simple

type Evolver interface {
	Evolve(Population, int, int) Population
}
