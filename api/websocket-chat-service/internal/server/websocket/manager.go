package websocket

import (
	"context"
	"time"
	"websocket-chat-service/init/logger"
	"websocket-chat-service/pkg/constants"

	"websocket-chat-service/internal/entities"
	"websocket-chat-service/internal/service"
)

type MessageManager interface {
	ManageMessage(ctx context.Context, message *entities.Message) error
}

type Manager struct {
	service service.ScyllaService
}

func NewManager(service service.ScyllaService) *Manager {
	return &Manager{
		service: service,
	}
}

func (m *Manager) ManageMessage(ctx context.Context, message *entities.Message) error {
	now := time.Now().Format("15:04:05 02-01-2006")
	message.SentAt = &now

	logger.Debug(message, constants.WebsocketCategory)

	return m.service.StoreMessage(ctx, message)
}
