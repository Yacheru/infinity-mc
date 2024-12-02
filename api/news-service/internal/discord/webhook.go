package discord

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/webhook"
	"news-service/init/logger"
	"news-service/pkg/constants"

	"news-service/init/config"
	"news-service/internal/entities"
)

type EmbedSender interface {
	SendEmbed(embed *entities.Embed) (int, error)
}

type WebhookClient struct {
	client webhook.Client
}

func NewWebhookClient(cfg *config.Config) *WebhookClient {
	return &WebhookClient{client: webhook.New(cfg.WebhookID, cfg.WebhookToken)}
}

func (w *WebhookClient) SendEmbed(embed *entities.Embed) (int, error) {
	embedBuilder := discord.NewEmbedBuilder()

	if embed.Title != "" {
		embedBuilder.SetTitle(embed.Title)
	}

	if embed.URL != "" {
		embedBuilder.SetURL(embed.URL)
	}

	if embed.Description != "" {
		embedBuilder.SetDescription(embed.Description)
	}

	if embed.CreatedAt != nil {
		embedBuilder.SetTimestamp(*embed.CreatedAt)
	}

	if embed.Image != nil {
		embedBuilder.SetImage(embed.Image.URL)
	}

	if embed.Thumbnail != nil {
		embedBuilder.SetThumbnail(embed.Thumbnail.URL)
	}

	if embed.Author != nil {
		embedBuilder.SetAuthor(embed.Author.Name, embed.Author.URL, embed.Author.IconURL)
	}

	if embed.Footer != nil {
		embedBuilder.SetFooter(embed.Footer.Text, embed.Footer.IconURL)
	}

	if embed.Fields != nil {
		fields := make([]discord.EmbedField, len(*embed.Fields))
		for i, field := range *embed.Fields {
			fields[i] = discord.EmbedField{
				Name:   field.Name,
				Value:  field.Value,
				Inline: field.Inline,
			}
		}
		embedBuilder.AddFields(fields...)
	}

	embedBuilder.SetColor(int(embed.Color))

	message, err := w.client.CreateEmbeds([]discord.Embed{embedBuilder.Build()})
	if err != nil {
		return 0, err
	}

	logger.DebugF("message sent successfully (ID: %d)", constants.LoggerDiscord, message.ID)

	return int(message.ID), nil
}
