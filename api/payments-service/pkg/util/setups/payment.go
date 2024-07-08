package setups

import (
	"fmt"
	records2 "payments-service/internal/records"

	"payments-service/internal/http/client"
)

func SetupPayment(handler *client.PaymentHandler, price, donat, nickname, duration, email string) (*records2.Payment, error) {
	payment, err := handler.CreatePayment(&records2.Payment{
		Amount: &records2.Amount{
			Value:    price,
			Currency: "RUB",
		},
		PaymentMethod: records2.PaymentMethodType("bank_card"),
		Receipt: &records2.Receipt{
			Customer: &records2.Email{
				Email: email,
			},
			Items: &[]records2.Items{
				{
					Description: fmt.Sprintf("Услуга %s", donat),
					Amount: &records2.Amount{
						Value:    price,
						Currency: "RUB",
					},
					VatCode:  1,
					Quantity: "1",
				},
			},
		},
		Metadata: &records2.Metadata{
			Nickname: nickname,
			Service:  donat,
			Price:    price,
			Duration: duration,
		},
		Capture: true,
		Confirmation: records2.Redirect{
			Type:      "redirect",
			ReturnURL: "https://infinity-mc.ru/",
		},
		Description: fmt.Sprintf("%s | %s | %s", nickname, donat, price),
	})

	return payment, err
}
