package redis

import (
	"rest/counter"

	"github.com/go-redis/redis"
)

type CounterRepository struct {
	Client *redis.Client
}

func NewCounterRepository(cli *redis.Client) *CounterRepository {
	return &CounterRepository{
		Client: cli,
	}
}

func (cr *CounterRepository) Add(i int) error {
	return cr.Client.IncrBy(counter.CounterKey, int64(i)).Err()
}
func (cr *CounterRepository) Sub(i int) error {
	return cr.Client.DecrBy(counter.CounterKey, int64(i)).Err()
}
func (cr *CounterRepository) Val() (int, error) {
	return cr.Client.Get(counter.CounterKey).Int()
}
