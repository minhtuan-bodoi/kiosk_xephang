package services

type QueueService struct{}

func NewQueueService() *QueueService {
	return &QueueService{}
}

func (s *QueueService) GetQueues() []string {
	return []string{}
}
