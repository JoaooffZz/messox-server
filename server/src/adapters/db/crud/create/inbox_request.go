package create

import (
	"fmt"
)

func (c *Create)NewInboxRequest(senderID int, adderessID int) error {
	query := fmt.Sprintf(`
	    INSERT INTO %s
		(sender_id, addressee_id)
		VALUES ($1, $2)
	`, tabelInboxRequests)

	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(query, senderID, adderessID)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}