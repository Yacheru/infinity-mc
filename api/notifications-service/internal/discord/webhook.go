package discord

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
	"github.com/sirupsen/logrus"

	"notifications-service/init/config"
	"notifications-service/init/logger"
	"notifications-service/internal/entities"
	"notifications-service/internal/utils"
	"notifications-service/pkg/constants"
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

	message, err := wh.CreateEmbeds([]discord.Embed{
		discord.NewEmbedBuilder().
			SetDescriptionf("Игрок **%s** купил **%s** на **%s**", s.Nickname, s.Service, s.Duration).
			SetColor(3140873).
			SetImage("https://imgur.com/a/8Hrf1Rc").
			Build(),
	})
	if err != nil {
		return err
	}

	logger.DebugF("message sent successfully (ID: %d)", logrus.Fields{constants.LoggerCategory: constants.DiscordCategory}, message.ID)

	return nil
}
