package create

import "fmt"

func (c *Create)NewContact(userID int, contactID int) error {

    query := fmt.Sprintf(`
	    INSERT INTO %s
		(user_id, contact_id)
		VALUES ($1, $2)
	`, tabelContacts)

	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err  = tx.Exec(query, userID, contactID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
	
}