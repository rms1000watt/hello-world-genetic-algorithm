package ga

type Individual interface {
	Fit() int // Rated from 0 - 100. 0 is worst, 100 is best
	Value() interface{}
	Breed(Individual) Individual
}
