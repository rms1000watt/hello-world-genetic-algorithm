package ga

// Grade is rated from 0 - 100
// where 0 is worst and 100 is best
type Population interface {
	Grade() int
	Sort() Population
	Best(int) Population
	Merge(Population) Population
	Length() int
	Mutate(int) Population
	At(int) Individual
	Push(Individual) Population
	Pop() (Population, Individual)
}
