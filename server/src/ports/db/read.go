package db

type ReadDB interface{
	GetTotalUsers() (int, error)
	GetUser(name string) (*User, error)
	GetUsers(name string) ([]ViewUser, error)
	GetInboxMessages(addresseeID int) ([]InboxMessage, error)
	GetSentRequests(senderID int) ([]ViewUser, error)
	GetReceivedRequests(addresseeID int) ([]ViewUser, error)
}