package repository

import (
	"log/slog"
	"xrChat_backend/config"
	"xrChat_backend/internal/model"
)

// SaveUnread save unread message.
func SaveUnread(message *model.Message) (err error) {
	err = config.DB.Create(message).Error
	if err != nil {
		slog.Error("save unread message failed", "err", err)
	}
	return
}
