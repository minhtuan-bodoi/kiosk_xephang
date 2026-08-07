package services

type CounterService struct {}

func NewCounterService() *CounterService {
    return &CounterService{}
}

func (s *CounterService) GetCounters() []string {
    return []string{}
}

