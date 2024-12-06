package service

import (
	"context"
	"encoding/json"
	"payments-service/init/config"
	"payments-service/init/logger"
	"payments-service/internal/entities"
	"payments-service/internal/kafka/producer"
	"payments-service/internal/repository"
	"payments-service/internal/server/http/client"
	"payments-service/internal/util"
	"payments-service/pkg/constants"
)

type Payments struct {
	repo     repository.PaymentsRepository
	producer producer.Producer
	cfg      *config.Config
}

func NewPayments(repo repository.PaymentsRepository) *Payments {
	return &Payments{repo: repo}
}

func (p *Payments) CreatePayment(ctx context.Context, payment *entities.PaymentService) (*entities.Payment, error) {
	yooClient := client.NewClient(p.cfg.YKassaID, p.cfg.YKassaPass)
	pH := client.NewPaymentHandler(yooClient)

	pay, err := util.SetupPayment(pH, payment)
	if err != nil {
		logger.Error(err.Error(), constants.LoggerCategoryService)
		return nil, err
	}

	return pay, nil
}

func (p *Payments) AcceptPayment(ctx context.Context, paid *entities.Paid) error {
	message, err := json.Marshal(&entities.KafkaMcMessage{
		Nickname: paid.Object.Metadata.Nickname,
		Duration: paid.Object.Metadata.Duration,
		Service:  paid.Object.Metadata.Service,
	})
	if err != nil {
		logger.Error(err.Error(), constants.LoggerCategoryService)
		return err
	}

	if err := p.producer.PrepareMessage(message); err != nil {
		logger.Error(err.Error(), constants.LoggerCategoryService)
		return err
	}

	if err := p.repo.StorePayment(ctx, paid.Object.Metadata); err != nil {
		logger.Error(err.Error(), constants.LoggerCategoryService)
		return err
	}

	return nil
}
