package delete

import (
	"fmt"
)

func (d *Delete)DeleteContact(userID int, contactID int) error {
	query := fmt.Sprintf(`
	    DELETE FROM %s
		WHERE user_id = $1
		AND contact_id = $2;
	`, tabelContacts)

	res, err := d.DB.Exec(query, userID, contactID)

    if err != nil {
        return err
    }

    rows, _ := res.RowsAffected()

    if rows == 0 {
        return fmt.Errorf("no lines deleted")
    }
	
    return nil
}