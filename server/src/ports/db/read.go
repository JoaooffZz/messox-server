package db

type ReadDB interface{
	TotalUsers() (int, error)
	Users(name string) ([]User, error)
	SentRequests(userID int64) ([]User, error)
	ReceivedRequests(userID int64) ([]User, error)
}