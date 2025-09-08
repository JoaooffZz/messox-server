package create

import "fmt"

const (
	contacts = "_contacts"
	inbox = "_inbox_message"
	chat = "_chat"
)

func (c *Create)TabelContacts(userID string) error {
	tabel := fmt.Sprintf("%s%s", userID, contacts)


	query := fmt.Sprintf(`
	    CREATE TABLE IF NOT EXISTS %s (
		    id_user BIGINT NOT NULL
		);
	`, tabel)

	_, err := c.DB.Exec(query)
    return err
}

func (c *Create)TabelInboxMessage(userID string) error {
	tabel := fmt.Sprintf("%s%s", userID, inbox)

   	query := fmt.Sprintf(`
	    CREATE TABLE IF NOT EXISTS %s (
		    id_sender BIGINT PRIMARY KEY NOT NULL,
			message TEXT NOT NULL,
			timestamp BIGINT NOT NULL
		);
	`, tabel)

	_, err := c.DB.Exec(query)
    return err
}

func (c *Create)TabelChat(chatID string) error {
	tabel := fmt.Sprintf("%s%s", chatID, chat)

   	query := fmt.Sprintf(`
	    CREATE TABLE IF NOT EXISTS %s (
		    date DATE PRIMARY KEY NOT NULL,
			chat JSONB NOT NULL
		);
	`, tabel)

	_, err := c.DB.Exec(query)
    return err
}