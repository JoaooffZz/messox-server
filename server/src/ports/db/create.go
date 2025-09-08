package db

type CreateDB interface {
	User(user User) (string, error)
	Chat(chat Chat) error
	Contact(userID string, contactID int64) error 
	InboxMessage(inbox InboxMessage) error
	InboxRequest(senderID int64, adderess int64) error

	TabelContacts(userID string) error
	TabelInboxMessage(userID string) error
	TabelChat(chatID string) error
}