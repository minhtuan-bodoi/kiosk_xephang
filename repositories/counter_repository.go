package repositories

type CounterRepository struct{}

func NewCounterRepository() *CounterRepository {
	return &CounterRepository{}
}
