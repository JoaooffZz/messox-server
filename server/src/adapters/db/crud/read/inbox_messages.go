package read

import (
	m "db/models"
	"fmt"
	"time"
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
		var createdAtStr string
		err := rows.Scan(&i.SenderID, &i.Message, &createdAtStr)
		if err != nil {
			return nil, err
		}
		i.CreatedAt, _ = time.Parse("15:04:05", createdAtStr)
		inboxs = append(inboxs, i)
    }

	return inboxs, nil
}