package telegram

type Telegram struct {
	Token string
}

func NewTelegram(token string) *Telegram {
	return &Telegram{Token: token}
}
