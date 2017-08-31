package simple

// Fit is rated from 0 - 100
// where 0 is worst and 100 is best
type Individual interface {
	Fit() int
	Value() interface{}
	Breed(Individual) Individual
}
