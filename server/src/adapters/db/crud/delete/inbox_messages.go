package delete

import "fmt"

func (d *Delete)DeleteInboxMessages(addresseeID int) error {
	query := fmt.Sprintf(`
	    DELETE ALL FROM %s
		WHERE addressee_id = $1;
	`, tabelInboxMessages)

	res, err := d.DB.Exec(query, addresseeID)

    if err != nil {
        return err
    }

    rows, _ := res.RowsAffected()

    if rows == 0 {
        return fmt.Errorf("no lines deleted")
    }
	
    return nil
}