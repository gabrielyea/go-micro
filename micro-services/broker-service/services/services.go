package services

import (
	"broker/listener"
	"context"
	"time"
)

type BrokerServiceInterface interface {
	Push(string, string, string) error
}

type brokerService struct {
	rbt listener.RabbitInterface
}

func NewBrokerService(rbt listener.RabbitInterface) BrokerServiceInterface {
	return &brokerService{rbt}
}

func (s *brokerService) Push(name, msg, key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	err := s.rbt.Push(name, msg, key, ctx)
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}
