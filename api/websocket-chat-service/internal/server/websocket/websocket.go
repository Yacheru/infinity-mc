package websocket

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/coder/websocket"
	"io"
	"sync"
	"websocket-chat-service/init/config"
	"websocket-chat-service/internal/service"

	"websocket-chat-service/init/logger"
	"websocket-chat-service/internal/entities"
	"websocket-chat-service/pkg/constants"
)

type Client interface {
	Dial(ctx context.Context) error
	CloseConnections() map[string]string
}

type WebSocket struct {
	url          string
	maxConnLimit int
	auth         string

	countConn int

	connections map[string]*websocket.Conn

	manager MessageManager

	wg *sync.WaitGroup
	mu sync.Mutex
}

func NewWebSocket(cfg *config.Config, service *service.Service) *WebSocket {
	manager := NewManager(service.ScyllaService)

	return &WebSocket{
		url:          cfg.WebsocketURL,
		maxConnLimit: cfg.WebsocketLimit,
		auth:         cfg.BearerAuth,
		connections:  make(map[string]*websocket.Conn, cfg.WebsocketLimit),
		manager:      manager,
		wg:           new(sync.WaitGroup),
		mu:           sync.Mutex{},
	}
}

func (ws *WebSocket) Dial(ctx context.Context) error {
	if len(ws.connections) == ws.maxConnLimit {
		return constants.MaxLimitConnError
	}

	c, _, err := websocket.Dial(ctx, ws.url, nil)
	if err != nil {
		logger.Error(err.Error(), "Dial: "+constants.WebsocketCategory)
		return err
	}

	if err := c.Write(ctx, websocket.MessageText, []byte(fmt.Sprintf("Bearer %s", ws.auth))); err != nil {
		logger.Error(err.Error(), "Write auth: "+constants.WebsocketCategory)
		c.Close(websocket.StatusTryAgainLater, "try again later")
		return err
	}
	if err := c.Write(ctx, websocket.MessageText, []byte("Listen PlayerChatEvent")); err != nil {
		logger.Error(err.Error(), "Write event: "+constants.WebsocketCategory)
		c.Close(websocket.StatusTryAgainLater, "try again later")
		return err
	}

	ws.mu.Lock()
	ws.countConn++
	ws.connections[fmt.Sprintf("Conn: #%d", ws.countConn)] = c
	ws.mu.Unlock()

	ws.listen(ctx, c)

	logger.Info("websocket connected", constants.WebsocketCategory)
	return nil
}

func (ws *WebSocket) CloseConnections() map[string]string {
	errorsMap := make(map[string]string, len(ws.connections))

	for key, conn := range ws.connections {
		err := conn.Close(websocket.StatusNormalClosure, "normal closure")
		if err != nil {
			logger.Error(err.Error(), "Close: "+constants.WebsocketCategory)
			errorsMap[key] = err.Error()
		}
	}

	ws.countConn = 0

	return errorsMap
}

func (ws *WebSocket) listen(ctx context.Context, c *websocket.Conn) {
	go func(c *websocket.Conn) {
		var message = new(entities.Message)
		for {
			select {
			case <-ctx.Done():
				logger.Info(ctx.Err().Error(), "Close websocket connection")
				return
			default:
				_, b, err := c.Read(ctx)
				if err != nil {
					if websocket.CloseStatus(err) != -1 {
						logger.Error("Websocket connection closed", "Read: "+constants.WebsocketCategory)
						return
					}
					if errors.Is(err, io.EOF) {
						logger.Error("Websocket connection closed by server", "Read: "+constants.WebsocketCategory)
						return
					}
					logger.Error(err.Error(), "WS Read: "+constants.WebsocketCategory)
					return
				}

				if b[0] != 'E' {
					continue
				}

				bytes := CutMessagePrefix(b)

				err = json.Unmarshal(bytes, message)
				if err != nil {
					logger.Error(err.Error(), constants.WebsocketCategory)
					continue
				}

				if err := ws.manager.ManageMessage(ctx, message); err != nil {
					continue
				}
			}
		}
	}(c)
}
