package db

type User struct {
	Name string
	Password string
	Profile string
	Bio string
}

type InboxMessage struct {
	SenderID string
	AdderessID string
	Message string
	Date int
}

type Chat struct {
	Date string
	Chat map[string]interface{}
}