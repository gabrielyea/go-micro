package services

type BrokerServiceInterface interface {
	Authenticate()
}

type brokerService struct {
}

func NewBrokerService() BrokerServiceInterface {
	return &brokerService{}
}

func (s *brokerService) Authenticate() {

}
