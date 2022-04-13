package counter

type ConterUseCase interface {
	Add(int) error
	Sub(int) error
	Val() (int, error)
}
