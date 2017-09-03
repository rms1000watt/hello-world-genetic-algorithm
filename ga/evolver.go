package ga

type Evolver interface {
	Evolve(Population, int, int, chan Individual) Population
}
