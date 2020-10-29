package notification

import (
	"notigram"
	"notigram/domain/repository"
)

type notificationUseCase struct {
	repository repository.NotificationRepository
}

func NewNotificationUseCase(repository repository.NotificationRepository) *notificationUseCase {
	return &notificationUseCase{repository: repository}
}

func (n *notificationUseCase) Notify(info *notigram.MessageInfo) error {
	n.repository.SendNews(info)
	return nil
}
