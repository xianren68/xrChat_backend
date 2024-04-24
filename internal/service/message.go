package service

import (
	"xrChat_backend/internal/model"
	"xrChat_backend/internal/proto/pb"
	"xrChat_backend/internal/repository"
)

// SaveUnread save message if target is offline.
func SaveUnread(msg *pb.Message, ty int) (err error) {
	message := &model.Message{}
	message.Src = uint(msg.Src)
	message.Tar = uint(msg.Tar)
	message.Msg = msg.Msg
	message.Type = ty
	return repository.SaveUnread(message)
}
