package models

type User struct {
	Name string
	Password string
	Profile string
	Bio string
}

type InboxMessage struct {
	AdderessID string
	SenderID int64
	Message string
	Timestamp int
}

type Chat struct {
	Date string
	Chat map[string]interface{}
}