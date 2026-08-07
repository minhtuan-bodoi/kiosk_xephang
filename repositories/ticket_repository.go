package repositories

type TicketRepository struct{}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{}
}
