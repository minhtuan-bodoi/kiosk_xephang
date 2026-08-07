package services

type ServiceService struct{}

func NewServiceService() *ServiceService {
	return &ServiceService{}
}

func (s *ServiceService) GetServices() []string {
	return []string{}
}
