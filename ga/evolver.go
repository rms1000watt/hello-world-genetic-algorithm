package ga

type Evolver interface {
	Evolve(Population, int, int, Migration) Population
}
