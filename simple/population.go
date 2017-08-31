package simple

type Population interface {
	Grade() int
	Sort() Population
	Best(int) Population
	Merge(Population) Population
	Length() int
	Mutate(int) Population
}
