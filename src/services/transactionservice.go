package services

import (
	"context"
	"github.com/NewChakrit/golang_gin_vuejs_full-stack/entity"
	"github.com/NewChakrit/golang_gin_vuejs_full-stack/repository"
)

type TransactionService interface {
	Add(ctx context.Context, transaction entity.Transaction) error
	Edit(ctx context.Context, transaction entity.Transaction) error
	Delete(ctx context.Context, id int64) error
	FindAll(ctx context.Context) ([]entity.Transaction, error)
}

type transactionService struct {
	repository repository.TransactionRepository
}

func NewTransactionService(repository repository.TransactionRepository) TransactionService {
	return &transactionService{
		repository: repository,
	}
}

func (s *transactionService) FindAll(ctx context.Context) ([]entity.Transaction, error) {
	transactions, err := s.repository.FindAll()
	return transactions, err
}

func (s *transactionService) Add(ctx context.Context, transaction entity.Transaction) error {
	err := s.repository.Add(transaction)
	return err
}

func (s *transactionService) Edit(ctx context.Context, transaction entity.Transaction) error {
	err := s.repository.Edit(ctx, transaction)
	return err
}

func (s *transactionService) Delete(ctx context.Context, id int64) error {
	err := s.repository.Delete(id)
	return err
}
