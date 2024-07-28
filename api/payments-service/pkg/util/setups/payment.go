package setups

import (
	"fmt"
	"payments-service/internal/entities"
	"payments-service/internal/server/http/client"
)

func SetupPayment(handler *client.PaymentHandler, price, donat, nickname, duration, email string) (*entities.Payment, error) {
	payment, err := handler.CreatePayment(&entities.Payment{
		Amount: &entities.Amount{
			Value:    price,
			Currency: "RUB",
		},
		PaymentMethod: entities.PaymentMethodType("bank_card"),
		Receipt: &entities.Receipt{
			Customer: &entities.Email{
				Email: email,
			},
			Items: &[]entities.Items{
				{
					Description: fmt.Sprintf("Услуга %s", donat),
					Amount: &entities.Amount{
						Value:    price,
						Currency: "RUB",
					},
					VatCode:  1,
					Quantity: "1",
				},
			},
		},
		Metadata: &entities.Metadata{
			Nickname: nickname,
			Service:  donat,
			Price:    price,
			Duration: duration,
		},
		Capture: true,
		Confirmation: entities.Redirect{
			Type:      "redirect",
			ReturnURL: "https://infinity-mc.ru/",
		},
		Description: fmt.Sprintf("%s | %s | %s", nickname, donat, price),
	})

	return payment, err
}
