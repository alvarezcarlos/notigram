package repository

import (
	"notigram"
)

type NotificationRepository interface {
	SendNews(info *notigram.MessageInfo) error
}