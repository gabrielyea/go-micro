package services

import (
	"context"
	"logger-service/models"
	"logger-service/repo"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type LoggerServInterface interface {
	Insert(models.LogEntry) error
	All() ([]*models.LogEntry, error)
	GetById(string) (*models.LogEntry, error)
	DropCollection() error
	UpdateCollection(*models.LogEntry) (*mongo.UpdateResult, error)
}

type loggerServices struct {
	r repo.LoggerRepoInterface
}

func NewLoggerServices(repo repo.LoggerRepoInterface) LoggerServInterface {
	return &loggerServices{repo}
}

func (s *loggerServices) Insert(entry models.LogEntry) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	err := s.r.Insert(ctx, entry)
	defer cancel()

	if err != nil {
		return err
	}
	return nil
}

func (s *loggerServices) All() ([]*models.LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	res, err := s.r.All(ctx)
	defer cancel()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *loggerServices) GetById(id string) (*models.LogEntry, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	res, err := s.r.GetById(ctx, id)
	defer cancel()

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *loggerServices) UpdateCollection(ref *models.LogEntry) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	res, err := s.r.UpdateCollection(ctx, ref)
	defer cancel()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *loggerServices) DropCollection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	err := s.r.DropCollection(ctx)
	defer cancel()

	if err != nil {
		return err
	}
	return nil
}
