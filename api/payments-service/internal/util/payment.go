package util

import (
	"fmt"

	"payments-service/internal/entities"
	"payments-service/internal/server/http/client"
)

func SetupPayment(handler *client.PaymentHandler, service *entities.PaymentService) (*entities.Payment, error) {
	payment, err := handler.CreatePayment(&entities.Payment{
		Amount: &entities.Amount{
			Value:    service.Price,
			Currency: "RUB",
		},
		PaymentMethod: entities.PaymentMethodType("bank_card"),
		Receipt: &entities.Receipt{
			Customer: &entities.Email{
				Email: service.Email,
			},
			Items: &[]entities.Items{
				{
					Description: fmt.Sprintf("Услуга %s", service.Service),
					Amount: &entities.Amount{
						Value:    service.Price,
						Currency: "RUB",
					},
					VatCode:  1,
					Quantity: "1",
				},
			},
		},
		Metadata: service,
		Capture:  true,
		Confirmation: entities.Redirect{
			Type:      "redirect",
			ReturnURL: "https://infinity-mc.ru/",
		},
		Description: fmt.Sprintf("%s | %s | %s", service.Nickname, service.Service, service.Price),
	})

	return payment, err
}
