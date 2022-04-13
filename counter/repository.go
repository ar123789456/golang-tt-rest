package counter

const CounterKey = "Counter"

type CounterRepository interface {
	Add(int) error
	Sub(int) error
	Val() (int, error)
}
