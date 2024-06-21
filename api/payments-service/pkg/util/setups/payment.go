package setups

import (
	"fmt"

	"payments-service/internal/http/client"
	"payments-service/internal/repository/records"
)

func SetupPayment(handler *client.PaymentHandler, price, donat, nickname, duration, email string) (*records.Payment, error) {
	payment, err := handler.CreatePayment(&records.Payment{
		Amount: &records.Amount{
			Value:    price,
			Currency: "RUB",
		},
		PaymentMethod: records.PaymentMethodType("bank_card"),
		Receipt: &records.Receipt{
			Customer: &records.Email{
				Email: email,
			},
			Items: &[]records.Items{
				{
					Description: fmt.Sprintf("Услуга %s", donat),
					Amount: &records.Amount{
						Value:    price,
						Currency: "RUB",
					},
					VatCode:  1,
					Quantity: "1",
				},
			},
		},
		Metadata: &records.Metadata{
			Nickname:  nickname,
			DonatType: donat,
			Price:     price,
			Duration:  duration,
		},
		Capture: true,
		Confirmation: records.Redirect{
			Type:      "redirect",
			ReturnURL: "https://infinity-mc.ru/",
		},
		Description: fmt.Sprintf("%s | %s | %s", nickname, donat, price),
	})

	return payment, err
}
