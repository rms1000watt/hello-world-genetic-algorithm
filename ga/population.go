package ga

type Population interface {
	Grade() int // Rated from 0 - 100. 0 is worst, 100 is best
	Sort() Population
	Best(int) Population
	Merge(Population) Population
	Length() int
	Mutate(int) Population
	At(int) Individual
	Push(Individual) Population
	Pop() (Individual, Population)
	Immigrate(int, Migration) Population
	Emigrate(int, Migration) Population
}
