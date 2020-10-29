package notigram

import (
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	notificationRepository "notigram/infrastructure/notification"
	"notigram/usecase/notification"
	"time"
)

func Notify(bot *tb.Bot, info *MessageInfo) {
	notificationRepo := notificationRepository.NewNotificationRepository(bot)
	notificationUseCase := notification.NewNotificationUseCase(notificationRepo)
	notificationUseCase.Notify(info)
}

func NewTelegramBot(url, token string) *tb.Bot {
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		URL:    url,
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}
	return b
}