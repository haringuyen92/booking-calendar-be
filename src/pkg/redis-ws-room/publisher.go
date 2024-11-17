package redis_ws_room

import (
	ws_room "booking-calendar-server-backend/internal/core/ws-room"
	"context"
	"encoding/json"
	"errors"
	"github.com/redis/go-redis/v9"
)

type WsPublisher interface {
	Publish(roomName string, message ws_room.Message) error
}

func NewWsPublisher(redis *redis.Client) WsPublisher {
	return &redisWsPublisher{
		redis: redis,
	}
}

type redisWsPublisher struct {
	redis *redis.Client
}

func (r redisWsPublisher) Publish(roomName string, message ws_room.Message) error {
	if r.redis == nil {
		return errors.New("redis client is nil")
	}
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return r.redis.Publish(
		context.Background(),
		roomName,
		data,
	).Err()
}
