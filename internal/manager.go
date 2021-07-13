package internal

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Manager struct {
	Token  string
	BotAPI *tgbotapi.BotAPI
}

func NewManager(token string) (*Manager, error) {
	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}
	return &Manager{
		Token:  token,
		BotAPI: botApi,
	}, nil
}

func (m *Manager) OnUpdate(token string, upd *tgbotapi.Update) error {
	if token != m.Token {
		return errors.New("invalid token")
	}
	_, err := m.BotAPI.Send(tgbotapi.MessageConfig{
		BaseChat:              tgbotapi.BaseChat{
			ChatID:              upd.Message.Chat.ID,
			ChannelUsername:     upd.Message.Chat.UserName,
		},
		Text:                  upd.Message.Text,
	})
	if err != nil {
		return err
	}
	return nil
}
