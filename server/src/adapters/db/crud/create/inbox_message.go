package create

import (
	"fmt"
	m "ports/db"
)

func (c *Create)NewInboxMessage(inbox m.InboxMessage) error {
	query := fmt.Sprintf(`
	    INSERT INTO %s
		(sender_id, addressee_id, message, created_at)
		VALUES ($1, $2, $3, $4)
	`, tabelInboxMessages)

	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(query, inbox.SenderID, inbox.AddresseeID, inbox.Message, inbox.CreatedAt)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}