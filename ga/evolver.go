package ga

type Evolver interface {
	Evolve(Population, int, int, int, Migration) Population
}
