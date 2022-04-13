package usecase

import "rest/counter"

type CounterUseCase struct {
	Repo counter.CounterRepository
}

func NewCounterUseCase(repo counter.CounterRepository) *CounterUseCase {
	return &CounterUseCase{
		Repo: repo,
	}
}

func (c *CounterUseCase) Add(i int) error {
	return c.Repo.Add(i)
}

func (c *CounterUseCase) Sub(i int) error {
	return c.Repo.Sub(i)
}

func (c *CounterUseCase) Val() (int, error) {
	return c.Repo.Val()
}
