package service

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"payments-service/init/config"
	"payments-service/internal/entities"
	producerMocks "payments-service/internal/kafka/producer/mocks"
	payMocks "payments-service/internal/repository/mocks"
	"testing"
)

func TestPayments_AcceptPayment(t *testing.T) {
	type args struct {
		ctx  context.Context
		paid *entities.Paid
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case #1",
			args: args{
				ctx: context.Background(),
				paid: &entities.Paid{
					Type:  "",
					Event: "",
					Object: &entities.Object{
						Metadata: &entities.PaymentService{
							Nickname: "yacheru",
							Email:    "yacheru@gmail.com",
							Service:  "hronon",
							Price:    "169",
							Duration: 156,
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			producer := producerMocks.NewProducer(t)
			repo := payMocks.NewPaymentsRepository(t)

			message, err := json.Marshal(&entities.KafkaMcMessage{
				Nickname: tt.args.paid.Object.Metadata.Nickname,
				Duration: tt.args.paid.Object.Metadata.Duration,
				Service:  tt.args.paid.Object.Metadata.Service,
			})
			if err != nil {
				t.Error(err)
			}

			producer.
				On("PrepareMessage", message).
				Return(nil).
				Once()

			repo.
				On("StorePayment", mock.Anything, tt.args.paid.Object.Metadata).
				Return(func(ctx context.Context, payment *entities.PaymentService) error {
					return nil
				}).
				Once()

			p := &Payments{
				repo:     repo,
				producer: producer,
				cfg:      &config.Config{},
			}
			if err := p.AcceptPayment(tt.args.ctx, tt.args.paid); (err != nil) != tt.wantErr {
				t.Errorf("AcceptPayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
