package db

type CreateDB interface {
	NewUser(user User) (int, error)
	NewHistoryChat(chat Chat) error
	NewContact(userID int, contactID int) error 
	NewInboxMessage(inbox InboxMessage) error
	NewInboxRequest(senderID int, adderessID int) error
}