package db

type DeleteDB interface{
	DelTabelChat(chatID string) error
	DelContact(userID string, contactID int64) error
}