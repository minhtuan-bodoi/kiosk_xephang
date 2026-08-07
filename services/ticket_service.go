package services

type TicketService struct{}

func NewTicketService() *TicketService {
	return &TicketService{}
}

func (s *TicketService) IssueTicket(queueID, serviceID string) string {
	return ""
}
