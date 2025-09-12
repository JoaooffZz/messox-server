package delete

import "database/sql"

const (
	tabelContacts = "contacts"
	tabelInboxMessages = "inbox_messages"
	tabelInboxRequests = "inbox_requests"
)

type Delete struct {
	DB *sql.DB
}