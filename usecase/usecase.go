package usecase

import (
	"notigram"
)

type NotificationUseCase interface {
	Notify(info *notigram.MessageInfo) error
}