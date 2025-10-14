package db

type DeleteDB interface{
	DeleteContact(userID int, contactID int) error
	DeleteInboxMessages(addresseeID int) error
	DeleteInboxRequest(senderID int, addresseeID int) error
}