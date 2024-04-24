package repository

import (
	"context"
	"xrChat_backend/config"
	"xrChat_backend/internal/model"
)

// AddUnRead add unread message to redis.
func AddUnRead(message *model.Message) (err error) {
	config.RedisClient.SAdd(context.Background(), "unRead", []uint{uint(message.Tar)})
	// add to mysql unread
	return SaveUnread(message)
}
