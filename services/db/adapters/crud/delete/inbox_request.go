package delete

import "fmt"

func (d *Delete)DeleteInboxRequest(senderID int, addresseeID int) error {
	query := fmt.Sprintf(`
	    DELETE FROM %s
		WHERE sender_id = $1
		AND addressee_id = $2;
	`, tabelInboxRequests)

	res, err := d.DB.Exec(query, senderID, addresseeID)

    if err != nil {
        return err
    }

    rows, _ := res.RowsAffected()

    if rows == 0 {
        return fmt.Errorf("no lines deleted")
    }
	
    return nil
}