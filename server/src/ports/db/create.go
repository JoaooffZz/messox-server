package db

type CreateDB interface {
	NewUser(name string, password string) (*User, error)
	NewHistoryChat(chat ChatData) error
	NewContact(userID int, contactID int) error 
	NewInboxMessage(inbox InboxMessage) error
	NewInboxRequest(senderID int, adderessID int) error
}