package db

type CreateDB interface {
	NewUser(user User) (int, error)
	NewHistoryChat(chat ChatData) error
	NewContact(userID int, contactID int) error 
	NewInboxMessage(inbox InboxMessage) error
	NewInboxRequest(senderID int, adderessID int) error
}