package ga

type Migration interface {
	Pop() (Individual, bool)
	Push(Individual) bool
	Flush()
}
