package discord

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
	"notifications-service/init/config"
	"notifications-service/internal/entities"
	"notifications-service/internal/utils"
)

type Sender interface {
	SendNotification(msg *entities.Message) error
}

type Webhook struct {
	webhook.Client
}

func NewWebhookClient(cfg *config.Config) *Webhook {
	client := webhook.New(cfg.WebhookID, cfg.WebhookToken)
	return &Webhook{client}
}

func (wh *Webhook) SendNotification(msg *entities.Message) error {
	s := utils.LocalizeStruct(msg)

	_, err := wh.CreateEmbeds([]discord.Embed{
		discord.NewEmbedBuilder().
			SetDescriptionf("Игрок **%s** купил **%s** на **%s**", s.Nickname, s.Service, s.Duration).
			SetColor(3140873).
			Build(),
	})
	if err != nil {
		return err
	}

	return nil
}
