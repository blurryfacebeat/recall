package users

import "time"

type User struct {
	ID               string    `json:"id"`
	TelegramID       int64     `json:"telegram_id,omitempty"`
	TelegramUsername string    `json:"telegram_username,omitempty"`
	DisplayName      string    `json:"display_name"`
	AvatarURL        string    `json:"avatar_url,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
