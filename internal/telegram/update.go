package telegram

type Update struct {
	UpdateID      int    `json:"update_id"`
	Message       string `json:"message"`
	EditedMessage string `json:"edited_message"`
}
