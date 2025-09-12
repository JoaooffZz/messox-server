package read

import (
	m "db/models"
	"fmt"
)

func (r *Read)GetInboxMessages(addresseeID int) ([]m.InboxMessage, error){
	query := fmt.Sprintf(`
	    SELECT 
		sender_id, message, created_at
		FROM %s
		WHERE addressee_id = $1;
	`, tabelInboxMessages)

	rows, err := r.DB.Query(query, addresseeID)

	if err != nil {
        return nil, err
    }
    defer rows.Close()

	var inboxs []m.InboxMessage

	for rows.Next() {
        var i m.InboxMessage
		err := rows.Scan(&i.SenderID, &i.Message, &i.CreatedAt)
		if err != nil {
			return nil, err
		}
		inboxs = append(inboxs, i)
    }

	return inboxs, nil
}