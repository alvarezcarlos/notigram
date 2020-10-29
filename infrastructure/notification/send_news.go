package notification

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"notigram"
)

type notificationRepository struct {
	bot *tb.Bot
}

func NewNotificationRepository(bot *tb.Bot) *notificationRepository {
	return &notificationRepository{bot: bot}
}

func (n *notificationRepository) SendNews(info *notigram.MessageInfo) error {
	chat, err := n.bot.ChatByID(info.ChatId)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	message := fmt.Sprintf("%s says: \n %s", info.AppName, info.Message)
	_, err = n.bot.Send(chat, message)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}