package rate

import (
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type Client struct {
	limiter *rate.Limiter
}

var client = make(map[string]*Client)
var mu sync.Mutex

func Limiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if client, exists := client[ip]; exists {
		return client.limiter
	}

	limiter := rate.NewLimiter(rate.Every(3*time.Second), 3)
	client[ip] = &Client{limiter}
	return limiter
}
