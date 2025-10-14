package create

import "database/sql"

const (
	tabelContacts = "contacts"
	tabelChats = "chats"
	tabelUsers = "users"
	tabelInboxMessages = "inbox_messages"
	tabelInboxRequests = "inbox_requests"
)

type Create struct {
	DB *sql.DB
}