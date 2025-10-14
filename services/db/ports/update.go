package db

import "time"

type UpdateDB interface{
	UpUserProfile(userID int, profile string) error
	UpUserBio(userID int, bio string) error
	UpHistoryChat(chatID int64, date time.Time, chat Chat) error
}